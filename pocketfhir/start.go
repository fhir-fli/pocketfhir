package pocketfhir

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

func StartPocketFHIR(dataDir string, hostname string, port string, getApiLogs bool, caddyPort string, caddyStoragePath string, certFile string, keyFile string) {
	// Set environment variables for PocketBase configuration
	log.Println("[DEBUG] Setting environment variables...")
	if err := os.Setenv("POCKETBASE_DATA_DIR", dataDir); err != nil {
		log.Fatalf("Failed to set data directory: %v", err)
	}

	// Handle graceful shutdown signal
	log.Println("[DEBUG] Setting up channel for graceful shutdown...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the PocketFHIR server in a separate goroutine
	go func() {
		log.Println("[DEBUG] Starting PocketFHIR server...")
		RunServer(dataDir, hostname, port, getApiLogs)
	}()

	// Start the Caddy server in a separate goroutine
	go func() {
		log.Println("[DEBUG] Starting Caddy server with HTTPS...")
		StartCaddy(caddyPort, fmt.Sprintf("http://%s:%s", hostname, port), caddyStoragePath, certFile, keyFile)
	}()

	// Wait for interrupt signal to gracefully shut down the server
	log.Println("[DEBUG] Waiting for interrupt signal to shut down the servers...")
	<-stop
	log.Println("Shutting down PocketFHIR and Caddy servers...")
	StopServer() // PocketFHIR shutdown
	log.Println("PocketFHIR and Caddy servers shut down gracefully.")
}
