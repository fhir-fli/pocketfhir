package pocketfhir

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caddyserver/caddy/v2"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

// CreateConfig generates a basic Caddy configuration to run a reverse proxy with HTTPS.
func CreateConfig(port, upstreamURL, storagePath, certFile, keyFile string) *caddy.Config {
	// Use storagePath for the log file paths
	caddyDebugLogPath := fmt.Sprintf("%s/caddy_debug.log", storagePath)
	accessLogPath := fmt.Sprintf("%s/access_log.log", storagePath)

	// JSON Configuration with dynamically inserted paths and certificates
	jsonConfig := fmt.Sprintf(`{
		"logging": {
			"logs": {
				"default": {
					"level": "DEBUG",
					"writer": {
						"output": "file",
						"filename": "%s"
					}
				},
				"http.access": {
					"level": "DEBUG",
					"writer": {
						"output": "file",
						"filename": "%s"
					},
					"encoder": {
						"format": "json"
					}
				}
			}
		},
		"apps": {
			"http": {
				"servers": {
					"srv0": {
						"listen": [
							":%s"
						],
						"routes": [
							{
								"handle": [
									{
										"handler": "request_body",
										"max_size": 10000000
									},
									{
										"handler": "reverse_proxy",
										"transport": {
											"protocol": "http",
											"read_timeout": 360000000000
										},
										"upstreams": [
											{
												"dial": "127.0.0.1:8090"
											}
										]
									}
								]
							}
						]
					}
				}
			}
		},
		"tls": {
			"certificates": {
				"automate": ["*"],
				"load_files": [
					{
						"certificate": "%s",
						"key": "%s"
					}
				]
			}
		}
	}`, caddyDebugLogPath, accessLogPath, port, certFile, keyFile)

	// Parse the JSON into a caddy.Config struct
	var caddyConfig caddy.Config
	err := json.Unmarshal([]byte(jsonConfig), &caddyConfig)
	if err != nil {
		log.Fatalf("Failed to parse JSON configuration: %v", err)
	}

	log.Printf("Generated Caddy Configuration: %s", jsonConfig)

	return &caddyConfig
}

// StartCaddy starts the Caddy server with the provided configuration.
func StartCaddy(port, upstreamURL, storagePath, certFile, keyFile string) {
	// Change working directory
	if err := os.Chdir(storagePath); err != nil {
		log.Fatalf("Failed to change working directory to %s: %v", storagePath, err)
	}

	// Log configuration for transparency
	log.Printf("Starting Caddy with configuration:\nPort: %s\nUpstreamURL: %s\nStoragePath: %s\nCertFile: %s\nKeyFile: %s\n", port, upstreamURL, storagePath, certFile, keyFile)

	// Generate Caddy config
	cfg := CreateConfig(port, upstreamURL, storagePath, certFile, keyFile)

	// Serialize for debugging
	configJSON, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalf("Failed to serialize Caddy config: %v", err)
	}
	log.Printf("Generated Caddy config: %s", string(configJSON))

	// Initialize Caddy
	log.Println("Initializing Caddy...")
	if err := caddy.Run(cfg); err != nil {
		log.Fatalf("Error running Caddy: %v", err)
	}

	log.Println("Caddy server started successfully.")
}
