package main

import "github.com/fhir-fli/pocketfhir/pocketfhir"

func main() {
	// Use the string wrapper for local development
	dataDir := "./assets"
	ipAddress := "127.0.0.1"
	port := "8080" // Changed from 443 to 8080 to avoid permission issues
	enableApiLogs := false

	pocketfhir.RunServer(dataDir, ipAddress, port, enableApiLogs)
}
