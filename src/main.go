package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// Standard setup for app
	standard(app)

	// Register hooks from hooks.go
	registerHooks(app)

	// Perform additional initialization after bootstrap
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		// Initialize pre-defined collections
		if err := initializePredefinedCollections(app); err != nil {
			return err
		}

		// Load search parameters from a JSON file
		searchParamsBundle, err := loadSearchParameters("search-parameters.json")
		if err != nil {
			return err
		}

		// Initialize collections based on search parameters
		if err := initializeCollections(app, searchParamsBundle); err != nil {
			return err
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
