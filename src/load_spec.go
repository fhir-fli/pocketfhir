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

// Load the FHIR spec (e.g., StructureDefinitions, CapabilityStatements) from NDJSON files
func loadFhirSpec(app *pocketbase.PocketBase) error {
	log.Println("Loading FHIR resources...")

	// Define paths to the NDJSON files
	ndjsonFiles := []string{
		"./fhir_definitions/capabilitystatement.ndjson",
		"./fhir_definitions/codesystem.ndjson",
		"./fhir_definitions/compartmentdefinition.ndjson",
		"./fhir_definitions/conceptmap.ndjson",
		"./fhir_definitions/operationdefinition.ndjson",
		"./fhir_definitions/searchparameter.ndjson",
		"./fhir_definitions/structuredefinition.ndjson",
		"./fhir_definitions/valueset.ndjson",
	}

	// Iterate over NDJSON file paths and load resources into the appropriate collections
	for _, ndjsonFilePath := range ndjsonFiles {
		collectionName := extractCollectionNameFromFile(ndjsonFilePath)
		if err := loadFhirResourcesFromNDJSON(app, ndjsonFilePath, collectionName); err != nil {
			return fmt.Errorf("failed to load resources from %s: %v", ndjsonFilePath, err)
		}
	}

	log.Println("FHIR resources loaded successfully.")
	return nil
}

// Extract the collection name from the NDJSON filename (e.g., 'codesystem.ndjson' -> 'codesystem')
func extractCollectionNameFromFile(filePath string) string {
	fileName := filepath.Base(filePath)
	return strings.TrimSuffix(fileName, ".ndjson")
}

// Load FHIR resources from an NDJSON file and insert them into the specified collection
func loadFhirResourcesFromNDJSON(app *pocketbase.PocketBase, filePath string, collectionName string) error {
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

	// Use bufio.Reader with a larger buffer
	reader := bufio.NewReader(file)

	for {
		// Read each line (which is a JSON object in NDJSON format)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break // End of file
			}
			return fmt.Errorf("error reading NDJSON file '%s': %v", filePath, err)
		}

		// Parse each line as a JSON object (FHIR resource)
		var resource map[string]interface{}
		if err := json.Unmarshal([]byte(line), &resource); err != nil {
			log.Printf("Failed to parse JSON from NDJSON file '%s': %v", filePath, err)
			continue
		}

		// Ensure the FHIR resource contains an "id" field
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
		record.Set("resource", types.JsonRaw([]byte(line)))

		// Save the record in the collection
		if err := app.Dao().SaveRecord(record); err != nil {
			log.Printf("Failed to save resource from NDJSON line in '%s': %v", filePath, err)
			continue
		}

		log.Printf("Successfully loaded resource with ID '%s' into collection '%s' from NDJSON file '%s'", resourceID, collectionName, filePath)
	}

	return nil
}
