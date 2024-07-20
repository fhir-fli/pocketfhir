package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
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

	updatedResourceBytes, err := updateResourceJson(resourceData, 1, resource.Id, resource.Updated.Time().UTC().Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("failed to update resource meta: %w", err)
	}

	resource.Set("resource", types.JsonRaw(updatedResourceBytes))
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

	updatedResourceBytes, err := updateResourceJson(resourceData, updatedVersionId, newResourceVersion.Id, newResourceVersion.Updated.Time().UTC().Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("failed to update resource meta: %w", err)
	}

	newResourceVersion.Set("resource", types.JsonRaw(updatedResourceBytes))
	newResourceVersion.Set("versionId", updatedVersionId)

	return nil
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
