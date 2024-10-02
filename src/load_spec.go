package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

// Load the FHIR spec (e.g., StructureDefinitions, CapabilityStatements) from JSON Bundle files only once
func loadFhirSpecOnce(app *pocketbase.PocketBase) error {
	// Check if the FHIR spec has already been loaded
	if isFhirSpecInitialized(app) {
		log.Println("FHIR resources already initialized, skipping.")
		return nil
	}

	log.Println("Loading FHIR resources for the first time...")

	// Define paths to the JSON files
	jsonFiles := []string{
		"./assets/fhir_spec/conceptmaps.json",
		"./assets/fhir_spec/dataelements.json",
		"./assets/fhir_spec/extension-definitions.json",
		"./assets/fhir_spec/profiles-others.json",
		"./assets/fhir_spec/profiles-resources.json",
		"./assets/fhir_spec/profiles-types.json",
		"./assets/fhir_spec/search-parameters.json",
		"./assets/fhir_spec/valuesets.json",
	}

	// Iterate over JSON file paths and load resources into the appropriate collections
	for _, jsonFilePath := range jsonFiles {
		collectionName := extractCollectionNameFromFile(jsonFilePath)
		if err := loadFhirResourcesFromJSON(app, jsonFilePath, collectionName); err != nil {
			return fmt.Errorf("failed to load resources from %s: %v", jsonFilePath, err)
		}
	}

	// Mark FHIR spec as initialized
	setFhirSpecInitialized(app)

	log.Println("FHIR resources loaded successfully.")
	return nil
}

// Load FHIR resources from a JSON Bundle file and insert them into the specified collection
func loadFhirResourcesFromJSON(app *pocketbase.PocketBase, filePath string, collectionName string) error {
	// Check if the collection exists
	collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
	if err != nil {
		return fmt.Errorf("collection '%s' not found: %v", collectionName, err)
	}

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

	// Ensure the file is a FHIR Bundle
	if fileData["resourceType"] != "Bundle" {
		return fmt.Errorf("file '%s' is not a valid FHIR Bundle", filePath)
	}

	// Process each resource entry in the Bundle
	entries, ok := fileData["entry"].([]interface{})
	if !ok {
		return fmt.Errorf("no 'entry' array found in Bundle file '%s'", filePath)
	}

	for _, entry := range entries {
		// Extract the 'resource' field from each entry in the Bundle
		resource, ok := entry.(map[string]interface{})["resource"].(map[string]interface{})
		if !ok {
			log.Printf("Invalid resource format in Bundle file '%s', skipping entry.", filePath)
			continue
		}

		// Ensure the resource contains an "id" field
		resourceID, ok := resource["id"].(string)
		if !ok || resourceID == "" {
			log.Printf("Resource does not contain a valid ID in '%s', skipping...", filePath)
			continue
		}

		// Create a new record for the resource in the specified collection
		record := models.NewRecord(collection)

		// Set the resource ID explicitly (this overrides PocketBase's random ID generation)
		record.Set("id", resourceID)

		// Set necessary fields (like resourceType and the resource itself)
		record.Set("resourceType", resource["resourceType"])
		resourceJSON, _ := json.Marshal(resource)
		record.Set("resource", types.JsonRaw(resourceJSON))

		// Save the record in the collection
		if err := app.Dao().SaveRecord(record); err != nil {
			log.Printf("Failed to save resource with ID '%s' from Bundle file '%s': %v", resourceID, filePath, err)
			continue
		}

		log.Printf("Successfully loaded resource with ID '%s' into collection '%s' from Bundle file '%s'", resourceID, collectionName, filePath)
	}

	return nil
}

// Extract the collection name from the JSON filename (e.g., 'codesystem.json' -> 'codesystem')
func extractCollectionNameFromFile(filePath string) string {
	fileName := filepath.Base(filePath)
	return strings.TrimSuffix(fileName, ".json")
}
