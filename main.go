package main

import "github.com/fhir-fli/pocketfhir/pocketfhir"

func main() {
	// Use the string wrapper for local development
	dataDir := "."
	pocketfhir.RunServer(dataDir)
}
