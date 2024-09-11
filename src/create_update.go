package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	"github.com/google/fhir/go/jsonformat/fhirvalidate"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
	"google.golang.org/protobuf/proto"
)

func handleResourceCreation(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	log.Println("Starting resource creation...")

	resource := e.Record // Directly access the Record from the event
	collectionName := resource.Collection().Name
	log.Printf("Collection name: %s", collectionName)

	// Skip history collections
	if strings.HasSuffix(collectionName, "history") {
		log.Println("Collection is a history collection. Skipping creation.")
		return nil
	}

	if !isVersionedCollection(app, collectionName) {
		log.Println("Collection is not versioned. Skipping.")
		return nil
	}

	resourceField := resource.Get("resource")
	resourceData, err := getResourceData(resourceField)
	if err != nil {
		log.Printf("Failed to get resource data: %v", err)
		return fmt.Errorf("failed to get resource data: %w", err)
	}

	// Validate the FHIR resource
	if err := validateFHIRResource(resourceData); err != nil {
		log.Printf("FHIR validation failed: %v", err)
		return fmt.Errorf("validation error: %w", err)
	}

	// Check and update meta fields (lastUpdated and versionId)
	resourceData, err = updateMetaFields(resourceData)
	if err != nil {
		log.Printf("Failed to update meta fields: %v", err)
		return fmt.Errorf("failed to update meta fields: %w", err)
	}

	// Re-serialize the resource and store it
	updatedResourceBytes, err := updateResourceJson(resourceData, resource.Id)
	if err != nil {
		log.Printf("Failed to update resource JSON: %v", err)
		return fmt.Errorf("failed to update resource meta: %w", err)
	}

	resource.Set("resource", types.JsonRaw(updatedResourceBytes))

	// Store FHIR data in PocketBase
	log.Println("Storing FHIR data...")
	if err := storeFHIRData(resourceData); err != nil {
		log.Printf("Failed to store FHIR data: %v", err)
		return fmt.Errorf("failed to store FHIR data: %w", err)
	}

	log.Println("Resource creation completed successfully.")
	return nil
}

func handleResourceUpdate(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	log.Println("Starting resource update...")

	newResourceVersion := e.Record // Directly access the Record from the event
	collectionName := newResourceVersion.Collection().Name
	log.Printf("Collection name: %s", collectionName)

	if !isVersionedCollection(app, collectionName) {
		log.Println("Collection is not versioned. Skipping.")
		return nil
	}

	resourceField := newResourceVersion.Get("resource")
	resourceData, err := getResourceData(resourceField)
	if err != nil {
		log.Printf("Failed to get resource data: %v", err)
		return fmt.Errorf("failed to get resource data: %w", err)
	}

	// Check if the meta.lastUpdated exists
	meta := getMetaFromResource(resourceData)
	if meta.LastUpdated == "" {
		return fmt.Errorf("Error: meta.lastUpdated is missing")
	}

	// Fetch the current resource from the main table
	currentResource, err := app.Dao().FindRecordById(collectionName, newResourceVersion.Id)
	if err != nil {
		log.Printf("Failed to fetch existing resource: %v", err)
		return fmt.Errorf("failed to fetch existing resource: %w", err)
	}

	// Compare the lastUpdated fields and move the older one to history
	currentResourceData := currentResource.Get("resource").([]byte)
	currentMeta := getMetaFromResource(currentResourceData)
	if currentMeta.LastUpdated > meta.LastUpdated {
		log.Println("New resource is older. Moving new resource to history.")
		if err := moveResourceToHistory(app, collectionName, newResourceVersion); err != nil {
			return err
		}
	} else {
		log.Println("New resource is newer. Moving current resource to history.")
		if err := moveResourceToHistory(app, collectionName, currentResource); err != nil {
			return err
		}

		// Update new resource version and save it in the main table
		updatedResourceBytes, err := updateResourceJson(resourceData, newResourceVersion.Id)
		if err != nil {
			log.Printf("Failed to update resource JSON: %v", err)
			return fmt.Errorf("failed to update resource meta: %w", err)
		}
		newResourceVersion.Set("resource", types.JsonRaw(updatedResourceBytes))

		log.Println("Storing updated FHIR data...")
		if err := storeFHIRData(resourceData); err != nil {
			log.Printf("Failed to store FHIR data: %v", err)
			return fmt.Errorf("failed to store FHIR data: %w", err)
		}
	}

	log.Println("Resource update completed successfully.")
	return nil
}

func updateResourceJson(resourceData []byte, id string) (updatedResourceBytes []byte, err error) {
	log.Println("Updating resource JSON...")

	var resourceJson map[string]interface{}
	if err := json.Unmarshal(resourceData, &resourceJson); err != nil {
		log.Printf("Failed to unmarshal resource JSON: %v", err)
		return nil, fmt.Errorf("failed to unmarshal resource JSON: %w", err)
	}

	resourceJson["id"] = id

	// Ensure that meta.lastUpdated and versionId are consistent
	meta, ok := resourceJson["meta"].(map[string]interface{})
	if !ok {
		meta = make(map[string]interface{})
	}

	if meta["lastUpdated"] != nil {
		lastUpdated := meta["lastUpdated"].(string)
		versionId := strings.ReplaceAll(lastUpdated, ":", ".")
		meta["versionId"] = versionId
		resourceJson["meta"] = meta
	}

	updatedResourceBytes, err = json.Marshal(resourceJson)
	if err != nil {
		log.Printf("Failed to marshal updated resource JSON: %v", err)
		return nil, fmt.Errorf("failed to marshal updated resource JSON: %w", err)
	}

	log.Println("Resource JSON updated successfully.")
	return updatedResourceBytes, nil
}

func isVersionedCollection(app *pocketbase.PocketBase, collectionName string) bool {
	collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
	if err != nil {
		log.Printf("Failed to find collection: %v", err)
		return false
	}

	for _, field := range collection.Schema.Fields() {
		if field.Name == "versionId" {
			return true
		}
	}
	return false
}

func updateMetaFields(resourceData []byte) ([]byte, error) {
	var resourceJson map[string]interface{}
	if err := json.Unmarshal(resourceData, &resourceJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource JSON: %w", err)
	}

	// Update the meta fields
	meta, ok := resourceJson["meta"].(map[string]interface{})
	if !ok {
		meta = make(map[string]interface{})
	}

	// Set lastUpdated if missing
	if meta["lastUpdated"] == nil {
		now := time.Now().UTC().Format(time.RFC3339)
		meta["lastUpdated"] = now
	}

	// Generate versionId based on lastUpdated
	lastUpdated := meta["lastUpdated"].(string)
	meta["versionId"] = strings.ReplaceAll(lastUpdated, ":", ".")

	resourceJson["meta"] = meta

	return json.Marshal(resourceJson)
}

func getResourceData(resourceField any) (resourceData []byte, err error) {
	log.Println("Getting resource data...")

	switch v := resourceField.(type) {
	case types.JsonRaw:
		resourceData = []byte(v)
	default:
		log.Println("Resource field is not of expected type")
		return nil, fmt.Errorf("resource field is not of expected type")
	}

	log.Println("Resource data retrieved successfully.")
	return resourceData, nil
}

func validateFHIRResource(resourceData []byte) error {
	log.Println("Validating FHIR resource...")

	unmarshaller, err := jsonformat.NewUnmarshaller("UTC", fhirversion.R4) // Use "UTC" as the time zone and fhirversion.R4 for the FHIR version
	if err != nil {
		log.Printf("Failed to create unmarshaller: %v", err)
		return fmt.Errorf("failed to create unmarshaller: %w", err)
	}

	msg, err := unmarshaller.Unmarshal(resourceData)
	if err != nil {
		log.Printf("Failed to unmarshal FHIR resource: %v", err)
		return fmt.Errorf("failed to unmarshal FHIR resource: %w", err)
	}

	if err := fhirvalidate.Validate(proto.Message(msg)); err != nil {
		log.Printf("FHIR resource validation failed: %v", err)
		return fmt.Errorf("FHIR resource validation failed: %w", err)
	}

	log.Println("FHIR resource validation successful.")
	return nil
}

func storeFHIRData(resourceData []byte) error {
	log.Println("Storing FHIR data...")

	var resourceJson map[string]interface{}
	if err := json.Unmarshal(resourceData, &resourceJson); err != nil {
		log.Printf("Failed to unmarshal resource JSON: %v", err)
		return fmt.Errorf("failed to unmarshal resource JSON: %w", err)
	}

	log.Printf("Successfully unmarshalled resource JSON")

	log.Printf("Successfully changed time zone")

	log.Println("FHIR data stored successfully.")
	return nil
}

func getMetaFromResource(resourceField any) *Meta {
	// Type assertion to ensure it's of type types.JsonRaw
	resourceData, ok := resourceField.(types.JsonRaw)
	if !ok {
		log.Println("Failed to assert resource field as types.JsonRaw")
		return &Meta{}
	}

	var resourceJson map[string]interface{}
	if err := json.Unmarshal([]byte(resourceData), &resourceJson); err != nil {
		log.Printf("Failed to unmarshal resource data: %v", err)
		return &Meta{}
	}

	// Extract meta from resource JSON
	meta, ok := resourceJson["meta"].(map[string]interface{})
	if !ok {
		log.Println("Meta field not found in resource")
		return &Meta{}
	}

	return &Meta{
		LastUpdated: meta["lastUpdated"].(string),
		VersionId:   meta["versionId"].(string),
	}
}

type Meta struct {
	LastUpdated string
	VersionId   string
}

func moveResourceToHistory(app *pocketbase.PocketBase, collectionName string, resource *models.Record) error {
	log.Printf("Moving resource with ID %s to history table...", resource.Id)

	// Define the history collection name
	historyCollectionName := collectionName + "_history"

	// Fetch the history collection
	historyCollection, err := app.Dao().FindCollectionByNameOrId(historyCollectionName)
	if err != nil {
		log.Printf("Failed to find history collection: %v", err)
		return fmt.Errorf("failed to find history collection: %w", err)
	}

	// Create a new historical record
	historicalResource := models.NewRecord(historyCollection)
	historicalResource.Set("fhirId", resource.Id)
	historicalResource.Set("resourceType", resource.Get("resourceType"))

	// Set the versionId and resource fields for the historical record
	meta := getMetaFromResource(resource.Get("resource"))
	historicalResource.Set("versionId", meta.LastUpdated) // Use lastUpdated as versionId for history
	historicalResource.Set("resource", resource.Get("resource"))

	// Save the historical record in the history table
	log.Printf("Saving historical resource with versionId %s to history table...", meta.VersionId)
	if err := app.Dao().SaveRecord(historicalResource); err != nil {
		log.Printf("Failed to save resource to history table: %v", err)
		return fmt.Errorf("failed to save resource to history table: %w", err)
	}

	log.Printf("Resource moved to history table successfully.")
	return nil
}
