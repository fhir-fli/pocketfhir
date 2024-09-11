package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	// Standard setup for app
	standard(app)

	// Register hooks from hooks.go
	registerHooks(app)

	// Initialize collections only if necessary
	if err := initializeCollections(app); err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	}

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start the app: %v", err)
	}
}
