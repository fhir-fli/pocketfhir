package main

import (
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

func createResource(app *pocketbase.PocketBase, resourceType string, resourceData []byte) error {
	collection, err := app.Dao().FindCollectionByNameOrId(resourceType)
	if err != nil {
		return fmt.Errorf("collection not found: %w", err)
	}

	record := models.NewRecord(collection)
	record.Set("resource", types.JsonRaw(resourceData))

	// Validate and process the resource
	if err := validateAndProcessResource(record, resourceData); err != nil {
		return err
	}

	return app.Dao().SaveRecord(record)
}

func readResource(app *pocketbase.PocketBase, resourceType string, id string) (*models.Record, error) {
	return app.Dao().FindRecordById(resourceType, id)
}

func updateResource(app *pocketbase.PocketBase, resourceType string, id string, resourceData []byte) error {
	record, err := app.Dao().FindRecordById(resourceType, id)
	if err != nil {
		return fmt.Errorf("resource not found: %w", err)
	}

	record.Set("resource", types.JsonRaw(resourceData))

	// Validate and process the resource
	if err := validateAndProcessResource(record, resourceData); err != nil {
		return err
	}

	return app.Dao().SaveRecord(record)
}

func deleteResource(app *pocketbase.PocketBase, resourceType string, id string) error {
	record, err := app.Dao().FindRecordById(resourceType, id)
	if err != nil {
		return fmt.Errorf("resource not found: %w", err)
	}

	return app.Dao().DeleteRecord(record)
}

func validateAndProcessResource(record *models.Record, resourceData []byte) error {
	// Validate the resource data
	if err := validateFHIRResource(resourceData); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Update the resource metadata (e.g., versionId and lastUpdated)
	updatedResourceBytes, err := updateResourceJson(resourceData, 1, record.Id, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("failed to update resource meta: %w", err)
	}

	record.Set("resource", types.JsonRaw(updatedResourceBytes))
	record.Set("versionId", 1)

	return nil
}
