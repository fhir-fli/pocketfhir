#!/bin/bash

# Navigate to the root directory of the project
cd "$(dirname "$0")/.."

# Clean up previous builds
rm -f pocketfhir_ios

# Initialize Go modules in the root directory if go.mod does not exist
if [ ! -f go.mod ]; then
    go mod init pocketfhir
fi

# Tidy up Go modules in the root directory
go mod tidy

# Build the PocketFHIR server for iOS (arm64 architecture)
echo "Building PocketFHIR for iOS..."
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o pocketfhir_ios ./src
echo "iOS PocketFHIR build completed."

# Start the PocketFHIR server in the background (for the regular platform)
./pocketfhir serve &

echo "PocketFHIR server started."
