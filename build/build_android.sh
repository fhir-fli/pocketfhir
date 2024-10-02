#!/bin/bash

# Navigate to the root directory of the project
cd "$(dirname "$0")/.."

# Clean up previous builds
rm -f pocketfhir_android

# Initialize Go modules in the root directory if go.mod does not exist
if [ ! -f go.mod ]; then
    go mod init pocketfhir
fi

# Tidy up Go modules in the root directory
go mod tidy

# Build the PocketFHIR server for Android and place the .aar package in the root directory
echo "Building PocketFHIR for Android..."
gomobile bind -androidapi 21 -o pocketfhir_android.aar ./src
echo "Android PocketFHIR build completed."
