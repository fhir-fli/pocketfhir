package main

import (
	"log"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema" // Import schema for SchemaField
	"github.com/pocketbase/pocketbase/tools/types"
)

func initializeCollections(app *pocketbase.PocketBase) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return createInitialCollections(app)
	})
	return nil
}

func createInitialCollections(app *pocketbase.PocketBase) error {
	if collections == nil {
		return nil
	}

	for _, collection := range collections {
		collectionName := strings.ToLower(collection["name"].(string))

		existingCollection, err := app.Dao().FindCollectionByNameOrId(collectionName)
		if err == nil && existingCollection != nil {
			log.Printf("Collection '%s' already exists. Skipping creation.", collectionName)
			continue
		}

		// Assuming schemaFields are inside the collection map itself
		schemaFields, ok := collection["schema"].([]map[string]interface{})
		if !ok {
			log.Printf("No schema fields found for collection '%s'. Skipping creation.", collectionName)
			continue
		}

		coll := &models.Collection{}
		form := forms.NewCollectionUpsert(app, coll)
		form.Name = collectionName
		form.Type = models.CollectionTypeBase
		form.ListRule = types.Pointer("@request.auth.id != ''")
		form.ViewRule = types.Pointer("@request.auth.id != ''")
		form.CreateRule = types.Pointer("@request.auth.id != ''")
		form.UpdateRule = types.Pointer("@request.auth.id != ''")
		form.DeleteRule = nil

		// Add fields to the schema based on the collection's schema definition
		for _, field := range schemaFields {
			form.Schema.AddField(&schema.SchemaField{
				Name:     field["name"].(string),
				Type:     field["type"].(string),
				Required: field["required"].(bool),
				Options:  field["options"].(map[string]interface{}),
			})
		}

		// Submit the form to create the collection
		if err := form.Submit(); err != nil {
			log.Printf("Failed to create collection '%s': %v", collectionName, err)
			continue
		}

		log.Printf("Collection '%s' created successfully", collectionName)
	}

	return nil
}
