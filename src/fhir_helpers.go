package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

// Check if the FHIR spec is already initialized
func isFhirSpecInitialized(app *pocketbase.PocketBase) bool {
	metadataCollection, err := app.Dao().FindCollectionByNameOrId("metadata")
	if err != nil {
		log.Printf("Metadata collection not found: %v", err)
		return false
	}

	// Check for a record indicating FHIR spec initialization
	record, err := app.Dao().FindFirstRecordByData(metadataCollection.Id, "fhirSpecInitialized", true)
	if err != nil || record == nil {
		return false
	}

	return true
}

// Mark the FHIR spec as initialized in the "metadata" collection
func setFhirSpecInitialized(app *pocketbase.PocketBase) {
	metadataCollection, err := app.Dao().FindCollectionByNameOrId("metadata")
	if err != nil {
		log.Printf("Metadata collection not found: %v", err)
		return
	}

	// Create a new record in the metadata collection
	record := models.NewRecord(metadataCollection)
	record.Set("fhirSpecInitialized", true)

	if err := app.Dao().SaveRecord(record); err != nil {
		log.Printf("Failed to save metadata record: %v", err)
	} else {
		log.Println("FHIR spec initialization marked as complete.")
	}
}
