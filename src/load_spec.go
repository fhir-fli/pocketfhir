package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

// Load the FHIR spec (e.g., StructureDefinitions, CapabilityStatements) from JSON files only once
func loadFhirSpecOnce(app *pocketbase.PocketBase) error {
	// Check if the FHIR spec has already been loaded
	if isFhirSpecInitialized(app) {
		log.Println("FHIR resources already initialized, skipping.")
		return nil
	}

	log.Println("Loading FHIR resources for the first time...")

	// Define paths to the JSON files
	jsonFiles := []string{
		"./assets/fhir_spec/capabilitystatement.json",
		"./assets/fhir_spec/codesystem.json",
		"./assets/fhir_spec/compartmentdefinition.json",
		"./assets/fhir_spec/conceptmap.json",
		"./assets/fhir_spec/operationdefinition.json",
		"./assets/fhir_spec/searchparameter.json",
		"./assets/fhir_spec/structuredefinition.json",
		"./assets/fhir_spec/valueset.json",
	}

	// Iterate over JSON file paths and load resources into the appropriate collections
	for _, jsonFilePath := range jsonFiles {
		if err := loadFhirResourcesFromJSON(app, jsonFilePath); err != nil {
			return fmt.Errorf("failed to load resources from %s: %v", jsonFilePath, err)
		}
	}

	// Mark FHIR spec as initialized
	setFhirSpecInitialized(app)

	log.Println("FHIR resources loaded successfully.")
	return nil
}

// Load FHIR resources from a JSON file and insert them into the corresponding collections
func loadFhirResourcesFromJSON(app *pocketbase.PocketBase, filePath string) error {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open JSON file '%s': %v", filePath, err)
	}
	defer file.Close()

	// Parse the JSON file
	var fileData map[string]interface{}
	if err := json.NewDecoder(file).Decode(&fileData); err != nil {
		return fmt.Errorf("failed to decode JSON file '%s': %v", filePath, err)
	}

	// Process each resource entry in the Bundle
	entries, ok := fileData["entry"].([]interface{})
	if !ok {
		return fmt.Errorf("no 'entry' array found in JSON file '%s'", filePath)
	}

	for _, entry := range entries {
		resource, ok := entry.(map[string]interface{})["resource"].(map[string]interface{})
		if !ok {
			log.Printf("Invalid resource format in JSON file '%s', skipping entry.", filePath)
			continue
		}

		// Determine the collection name based on the resourceType (converted to lowercase)
		resourceType, ok := resource["resourceType"].(string)
		if !ok || resourceType == "" {
			log.Printf("Resource missing 'resourceType', skipping entry.")
			continue
		}
		collectionName := strings.ToLower(resourceType)

		// Ensure the resource contains an "id" field
		resourceID, ok := resource["id"].(string)
		if !ok || resourceID == "" {
			log.Printf("Resource does not contain a valid ID, skipping entry.")
			continue
		}

		// Check if the collection exists
		collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
		if err != nil {
			log.Printf("Collection '%s' not found for resourceType '%s', skipping entry.", collectionName, resourceType)
			continue
		}

		// Create a new record for the resource in the corresponding collection
		record := models.NewRecord(collection)
		record.Set("id", resourceID) // Set the resource ID
		resourceJSON, _ := json.Marshal(resource)
		record.Set("resource", types.JsonRaw(resourceJSON)) // Set the resource JSON

		// Save the record in the collection
		if err := app.Dao().SaveRecord(record); err != nil {
			log.Printf("Failed to save resource with ID '%s' to collection '%s': %v", resourceID, collectionName, err)
			continue
		}

		log.Printf("Successfully loaded resource with ID '%s' into collection '%s' from JSON file '%s'", resourceID, collectionName, filePath)
	}

	return nil
}
