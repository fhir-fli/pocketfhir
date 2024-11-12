package pocketfhir

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

// RunServer initializes and runs the PocketFHIR server
func RunServer() {
	app := pocketbase.New()

	// Standard setup for app
	standard(app)

	// Register hooks from hooks.go
	registerHooks(app)

	// Register FHIR routes
	registerFHIRRoutes(app)

	// Initialize collections only if necessary
	if err := initializeCollections(app); err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	}

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start the app: %v", err)
	}
}
