package pocketfhir

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caddyserver/caddy/v2"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

// CreateConfig generates a basic Caddy configuration to run a reverse proxy.
func CreateConfig(port, upstreamURL, storagePath string) *caddy.Config {
	return &caddy.Config{
		Admin: &caddy.AdminConfig{
			Listen: fmt.Sprintf(":%s", port),
		},
		StorageRaw: json.RawMessage(fmt.Sprintf(`{
			"module": "file_system",
			"root": "%s"
		}`, storagePath)),
		AppsRaw: map[string]json.RawMessage{
			"http": json.RawMessage(fmt.Sprintf(`{
				"servers": {
					"simple_reverse_proxy": {
						"listen": [":%s"],
						"routes": [{
							"handle": [{
								"handler": "reverse_proxy",
								"upstreams": [{
									"dial": "%s"
								}]
							}]
						}]
					}
				}
			}`, port, upstreamURL)), // Ensure upstreamURL is used correctly
		},
	}
}

func StartCaddy(port, upstreamURL, storagePath string) {
	if err := os.Chdir(storagePath); err != nil {
		log.Fatalf("Failed to change working directory to %s: %v", storagePath, err)
	}

	log.Printf("Starting Caddy with configuration:\nPort: %s\nUpstreamURL: %s\nStoragePath: %s\n", port, upstreamURL, storagePath)

	cfg := CreateConfig(port, upstreamURL, storagePath)
	configJSON, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalf("Failed to serialize Caddy config: %v", err)
	}
	log.Printf("Generated Caddy config: %s", string(configJSON))

	// Start the Caddy server (no blocking)
	log.Println("Initializing Caddy...")
	if err := caddy.Run(cfg); err != nil {
		log.Fatalf("Error running Caddy: %v", err)
	}

	log.Println("Caddy server started successfully.")
}
