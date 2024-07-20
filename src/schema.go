package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func createTableForResource(app *pocketbase.PocketBase, resourceName string, searchParams []SearchParameter) error {
	collection := &models.Collection{}
	collection.Name = resourceName
	form := forms.NewCollectionUpsert(app, collection)
	form.Name = resourceName
	form.Type = models.CollectionTypeBase
	form.ListRule = types.Pointer("@request.auth.id != ''")
	form.ViewRule = types.Pointer("@request.auth.id != ''")
	form.CreateRule = types.Pointer("@request.auth.id != ''")
	form.UpdateRule = types.Pointer("@request.auth.id != ''")
	form.DeleteRule = nil

	for _, param := range searchParams {
		field := &schema.SchemaField{
			Name:     param.Code,
			Type:     getFieldType(param.Type),
			Required: false,
			Options:  map[string]interface{}{},
		}
		form.Schema.AddField(field)
	}

	return form.Submit()
}

func getFieldType(paramType string) string {
	switch paramType {
	case "string":
		return schema.FieldTypeText
	case "token":
		return schema.FieldTypeText
	case "date":
		return schema.FieldTypeDate
	case "reference":
		return schema.FieldTypeRelation
	case "boolean":
		return schema.FieldTypeBool
	default:
		return schema.FieldTypeJson
	}
}

func initializeCollections(app *pocketbase.PocketBase, searchParamsBundle *Bundle) error {
	for _, entry := range searchParamsBundle.Entry {
		searchParam := entry.Resource
		for _, base := range searchParam.Base {
			if err := createTableForResource(app, base, []SearchParameter{searchParam}); err != nil {
				log.Printf("Failed to create table for resource '%s': %v", base, err)
				continue
			}
			log.Printf("Table for resource '%s' created successfully", base)
		}
	}
	return nil
}
