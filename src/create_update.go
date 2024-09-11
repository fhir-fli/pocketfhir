package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fhir-fli/fhirpath-go/fhir"
	"github.com/fhir-fli/fhirpath-go/pkg/containedresource"
	"github.com/fhir-fli/fhirpath-go/pkg/fhirwrapper"
	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	"github.com/google/fhir/go/jsonformat/fhirvalidate"
	bcrpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
	"google.golang.org/protobuf/proto"
)

func handleResourceCreation(app *pocketbase.PocketBase, e *core.ModelEvent) error {
	resource, ok := e.Model.(*models.Record)
	if !ok {
		return nil
	}

	collectionName := resource.Collection().Name
	if strings.HasSuffix(collectionName, "history") {
		return nil
	}

	if !isVersionedCollection(app, collectionName) {
		return nil
	}

	resourceField := resource.Get("resource")
	resourceData, err := getResourceData(resourceField)
	if err != nil {
		return fmt.Errorf("failed to get resource data: %w", err)
	}

	if err := validateFHIRResource(resourceData); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	searchParams := getSearchParamsForResource(collectionName)
	results, err := evaluateFHIRPathExpressions(resourceData, searchParams)
	if err != nil {
		return fmt.Errorf("failed to evaluate FHIRPath expressions: %w", err)
	}

	for k, v := range results {
		if v != nil {
			resource.Set(k, v)
		} else {
			resource.Set(k, nil) // Ensure that optional fields are explicitly set to null if they are empty
		}
	}

	updatedResourceBytes, err := updateResourceJson(resourceData, 1, resource.Id, resource.Updated.Time().UTC().Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("failed to update resource meta: %w", err)
	}

	resource.Set("resource", types.JsonRaw(updatedResourceBytes))

	if err := storeFHIRData(resource, resourceData); err != nil {
		return fmt.Errorf("failed to store FHIR data: %w", err)
	}

	resource.Set("versionId", 1)

	return nil
}

func handleResourceUpdate(app *pocketbase.PocketBase, e *core.ModelEvent) error {
	newResourceVersion, ok := e.Model.(*models.Record)
	if !ok {
		return nil
	}

	collectionName := newResourceVersion.Collection().Name
	if !isVersionedCollection(app, collectionName) {
		return nil
	}

	resourceField := newResourceVersion.Get("resource")
	resourceData, err := getResourceData(resourceField)
	if err != nil {
		return fmt.Errorf("failed to get resource data: %w", err)
	}

	if err := validateFHIRResource(resourceData); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	currentResourceVersion, err := app.Dao().FindRecordById(collectionName, newResourceVersion.Id)
	if err != nil {
		return fmt.Errorf("failed to fetch existing resource: %w", err)
	}

	historyCollectionName := collectionName + "history"
	historyCollection, err := app.Dao().FindCollectionByNameOrId(historyCollectionName)
	if err != nil {
		return fmt.Errorf("failed to create history collection: %w", err)
	}

	historicalResourceVersion := models.NewRecord(historyCollection)
	historicalResourceVersion.Set("fhirId", currentResourceVersion.Id)
	historicalResourceVersion.Set("resourceType", currentResourceVersion.Get("resourceType"))
	historicalResourceVersionId := currentResourceVersion.GetInt("versionId")
	updatedVersionId := historicalResourceVersionId + 1
	historicalResourceVersion.Set("versionId", historicalResourceVersionId)
	historicalResourceVersion.Set("resource", currentResourceVersion.Get("resource"))

	saveErr := app.Dao().SaveRecord(historicalResourceVersion)
	if saveErr != nil {
		return fmt.Errorf("failed to save resource to history table: %w", saveErr)
	}

	searchParams := getSearchParamsForResource(collectionName)
	results, err := evaluateFHIRPathExpressions(resourceData, searchParams)
	if err != nil {
		return fmt.Errorf("failed to evaluate FHIRPath expressions: %w", err)
	}

	for k, v := range results {
		if v != nil {
			newResourceVersion.Set(k, v)
		} else {
			newResourceVersion.Set(k, nil) // Ensure that optional fields are explicitly set to null if they are empty
		}
	}

	updatedResourceBytes, err := updateResourceJson(resourceData, updatedVersionId, newResourceVersion.Id, newResourceVersion.Updated.Time().UTC().Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("failed to update resource meta: %w", err)
	}

	newResourceVersion.Set("resource", types.JsonRaw(updatedResourceBytes))

	if err := storeFHIRData(newResourceVersion, resourceData); err != nil {
		return fmt.Errorf("failed to store FHIR data: %w", err)
	}

	newResourceVersion.Set("versionId", updatedVersionId)

	return nil
}

func getSearchParamsForResource(resourceType string) []SearchParameter {
	if params, exists := searchParamsByResourceType[resourceType]; exists {
		return params
	}
	return []SearchParameter{}
}

func updateResourceJson(resourceData []byte, newVersionId int, id string, lastUpdated string) (updatedResourceBytes []byte, err error) {
	var resourceJson map[string]interface{}
	if err := json.Unmarshal(resourceData, &resourceJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource JSON: %w", err)
	}

	resourceJson["id"] = id

	if resourceJson["meta"] == nil {
		resourceJson["meta"] = map[string]interface{}{}
	}

	meta := resourceJson["meta"].(map[string]interface{})
	meta["versionId"] = fmt.Sprintf("%d", newVersionId)
	meta["lastUpdated"] = lastUpdated
	return json.Marshal(resourceJson)
}

func isVersionedCollection(app *pocketbase.PocketBase, collectionName string) bool {
	collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
	if err != nil {
		return false
	}

	for _, field := range collection.Schema.Fields() {
		if field.Name == "versionId" {
			return true
		}
	}
	return false
}

func getResourceData(resourceField any) (resourceData []byte, err error) {
	switch v := resourceField.(type) {
	case types.JsonRaw:
		resourceData = []byte(v)
	default:
		return nil, fmt.Errorf("resource field is not of expected type")
	}
	return resourceData, nil
}

func validateFHIRResource(resourceData []byte) error {
	unmarshaller, err := jsonformat.NewUnmarshaller("UTC", fhirversion.R4)
	if err != nil {
		return fmt.Errorf("failed to create unmarshaller: %w", err)
	}

	msg, err := unmarshaller.Unmarshal(resourceData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal FHIR resource: %w", err)
	}

	if err := fhirvalidate.Validate(proto.Message(msg)); err != nil {
		return fmt.Errorf("FHIR resource validation failed: %w", err)
	}

	return nil
}

func storeFHIRData(resource *models.Record, resourceData []byte) error {
	var resourceJson map[string]interface{}
	if err := json.Unmarshal(resourceData, &resourceJson); err != nil {
		return fmt.Errorf("failed to unmarshal resource JSON: %w", err)
	}

	unmarshaller, err := jsonformat.NewUnmarshaller("r4", fhirversion.R4)
	if err != nil {
		return fmt.Errorf("failed to create unmarshaller: %w", err)
	}

	unmarshalledResource, err := unmarshaller.Unmarshal(resourceData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal resource to proto: %w", err)
	}

	containedResource := containedresource.Wrap(unmarshalledResource.(fhir.Resource))

	searchParams := getSearchParamsForResource(resource.Collection().Name)
	for _, param := range searchParams {
		value, err := evaluateFHIRPath(containedResource, param.Expression)
		if err != nil {
			return fmt.Errorf("failed to evaluate FHIRPath expression %s: %w", param.Expression, err)
		}

		if len(value) > 0 {
			resource.Set(param.Code, value[0])
		}
	}

	return nil
}

func evaluateFHIRPath(resource *bcrpb.ContainedResource, expression string) ([]interface{}, error) {
	compiledExpr, err := fhirwrapper.CompileFHIRPath(expression)
	if err != nil {
		return nil, fmt.Errorf("failed to compile FHIRPath expression: %w", err)
	}
	result, err := fhirwrapper.EvaluateFHIRPath(compiledExpr, resource)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate FHIRPath expression: %w", err)
	}
	return result, nil
}
