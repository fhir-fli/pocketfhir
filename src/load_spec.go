package main

import (
	"bufio"
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

// Load MIMIC-IV data from NDJSON files and insert into the appropriate collections
func loadMimicIVData(app *pocketbase.PocketBase) error {
	log.Println("Loading MIMIC-IV dataset...")

	// Directory containing the MIMIC-IV NDJSON files
	mimicIVDir := "./assets/mimic_iv/"

	// List NDJSON files in the directory
	files, err := os.ReadDir(mimicIVDir)
	if err != nil {
		return fmt.Errorf("failed to read MIMIC-IV directory: %v", err)
	}

	// Process each NDJSON file
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".ndjson") {
			continue // Skip non-NDJSON files
		}

		// Get the collection name from the file name (minus the .ndjson extension)
		collectionName := strings.TrimSuffix(file.Name(), ".ndjson")

		// Process the NDJSON file and load the resources
		ndjsonFilePath := filepath.Join(mimicIVDir, file.Name())
		if err := loadNDJSONFile(app, ndjsonFilePath, collectionName); err != nil {
			return fmt.Errorf("failed to load resources from %s: %v", ndjsonFilePath, err)
		}
	}

	log.Println("MIMIC-IV dataset loaded successfully.")
	return nil
}

// Load resources from a specific NDJSON file and insert them into the specified collection
func loadNDJSONFile(app *pocketbase.PocketBase, filePath string, collectionName string) error {
	// Check if the collection exists
	collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
	if err != nil {
		return fmt.Errorf("collection '%s' not found: %v", collectionName, err)
	}

	// Open the NDJSON file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open NDJSON file '%s': %v", filePath, err)
	}
	defer file.Close()

	// Initialize scanner to read NDJSON file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse the JSON line into a map
		var resource map[string]interface{}
		if err := json.Unmarshal([]byte(line), &resource); err != nil {
			log.Printf("Failed to parse JSON line in '%s': %v", filePath, err)
			continue
		}

		// Ensure the resource contains an "id" field (or any other relevant key for MIMIC-IV resources)
		resourceID, ok := resource["id"].(string)
		if !ok || resourceID == "" {
			log.Printf("Resource does not contain a valid 'id' in '%s', skipping entry.", filePath)
			continue
		}

		// Create a new record for the resource in the specified collection
		record := models.NewRecord(collection)
		record.Set("id", resourceID)
		resourceJSON, _ := json.Marshal(resource)
		record.Set("resource", types.JsonRaw(resourceJSON))

		// Save the record in the collection
		if err := app.Dao().SaveRecord(record); err != nil {
			log.Printf("Failed to save resource with ID '%s' from file '%s': %v", resourceID, filePath, err)
			continue
		}

		log.Printf("Successfully loaded resource with ID '%s' into collection '%s' from file '%s'", resourceID, collectionName, filePath)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading NDJSON file '%s': %v", filePath, err)
	}

	return nil
}
