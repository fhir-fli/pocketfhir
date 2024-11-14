package pocketfhir

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pocketbase/pocketbase"
)

func RunServer(dataDir string) {
	// Set the environment variable for Pocketbase data directory
	if err := os.Setenv("POCKETBASE_DATA_DIR", dataDir); err != nil {
		log.Fatalf("Failed to set data directory: %v", err)
	}

	app := pocketbase.New()

	// Standard setup for app
	standard(app)

	// Register hooks from hooks.go
	registerHooks(app)

	// Register FHIR routes
	registerFHIRRoutes(app)

	// Register server management routes
	registerManagementRoutes(app)

	// Initialize collections only if necessary
	if err := initializeCollections(app); err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	}

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Start(); err != nil {
			log.Fatalf("Failed to start the app: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	<-stop
	log.Println("Shutting down PocketFHIR server...")
	app.ResetBootstrapState() // Optional: Clean up any resources before shutting down
}
