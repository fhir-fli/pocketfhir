package main

import (
	"log"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var collections = []map[string]interface{}{
	{"name": "Account", "schema": generateSchemaFields("Account")},
	// Add more resources as needed
}

func initializeCollections(app *pocketbase.PocketBase) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return createInitialCollections(app)
	})
	return nil
}

func createInitialCollections(app *pocketbase.PocketBase) error {
	db := app.Dao().DB()

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

		schemaFields := generateSchemaFields(collectionName)
		if schemaFields == nil {
			log.Printf("No schema fields found for collection '%s'. Skipping creation.", collectionName)
			continue
		}

		indexes := indexesByResourceType[collectionName]
		coll := &models.Collection{}
		form := forms.NewCollectionUpsert(app, coll)
		form.Name = collectionName
		form.Type = models.CollectionTypeBase
		form.ListRule = types.Pointer("@request.auth.id != ''")
		form.ViewRule = types.Pointer("@request.auth.id != ''")
		form.CreateRule = types.Pointer("@request.auth.id != ''")
		form.UpdateRule = types.Pointer("@request.auth.id != ''")
		form.DeleteRule = nil

		for _, field := range schemaFields {
			form.Schema.AddField(&models.SchemaField{
				Name:     field["name"].(string),
				Type:     field["type"].(string),
				Required: field["required"].(bool),
				Options:  field["options"].(map[string]interface{}),
			})
		}

		if err := form.Submit(); err != nil {
			log.Printf("Failed to create collection '%s': %v", collectionName, err)
			continue
		}

		for _, index := range indexes {
			indexName := index["name"].(string)
			indexType := index["type"].(string)
			fields := strings.Join(index["fields"].([]string), ", ")

			query := "CREATE " + indexType + " INDEX " + indexName + " ON " + collectionName + " (" + fields + ");"
			if _, err := db.NewQuery(query).Execute(); err != nil {
				log.Printf("Failed to create index '%s' on collection '%s': %v", indexName, collectionName, err)
			} else {
				log.Printf("Index '%s' created successfully on collection '%s'", indexName, collectionName)
			}
		}

		log.Printf("Collection '%s' created successfully", collectionName)
	}

	return nil
}
