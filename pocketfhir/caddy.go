package pocketfhir

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caddyserver/caddy/v2"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

// StartCaddy starts the Caddy server with the provided configuration.
func StartCaddy(pbPort, httpPort, httpsPort, pbUrl, storagePath, ipAddress string) {
	// Change working directory
	if err := os.Chdir(storagePath); err != nil {
		log.Fatalf("Failed to change working directory to %s: %v", storagePath, err)
	}

	// Log configuration for transparency
	log.Printf("Starting Caddy server with the following configuration:")
	log.Printf("PocketBase Port: %s", pbPort)
	log.Printf("HTTP Port: %s", httpPort)
	log.Printf("HTTPS Port: %s", httpsPort)
	log.Printf("Upstream URL: %s", pbUrl)
	log.Printf("Storage Path: %s", storagePath)

	// Generate Caddy config
	cfg := createConfig(pbPort, httpPort, httpsPort, pbUrl, storagePath, ipAddress)

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

// CreateConfig generates a basic Caddy configuration to run a reverse proxy with HTTP and a static file server.
func createConfig(pbPort, httpPort, httpsPort, pbUrl, storagePath, ipAddress string) *caddy.Config {
	// Use storagePath for the log file paths
	caddyDebugLogPath := fmt.Sprintf("%s/caddy_debug.log", storagePath)
	caddyHTTPSDebugLogPath := fmt.Sprintf("%s/caddy_https_debug.log", storagePath)
	rootCertPath := fmt.Sprintf("%s/pki/authorities/local", storagePath)

	// Generate the JSON configuration
	jsonConfig := jsonConfig(pbPort, httpPort, httpsPort, pbUrl, storagePath, ipAddress, caddyDebugLogPath, caddyHTTPSDebugLogPath, rootCertPath)
	log.Printf("Generated JSON Configuration: %s", jsonConfig)

	// Parse the JSON into a caddy.Config struct
	var caddyConfig caddy.Config
	err := json.Unmarshal([]byte(jsonConfig), &caddyConfig)
	if err != nil {
		log.Fatalf("Failed to parse JSON configuration: %v", err)
	}

	log.Printf("Generated Caddy Configuration: %s", jsonConfig)

	return &caddyConfig
}

func jsonConfig(pbPort, httpPort, httpsPort, pbUrl, storagePath, ipAddress, caddyDebugLogPath, caddyHTTPSDebugLogPath, rootCertPath string) string {
	// JSON Configuration for Caddy with a basic file server
	return fmt.Sprintf(`{
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
        "storage": {
            "module": "file_system",
            "root": "%s"
        },
        "apps": {
            "http": {
                "http_port": %s,
                "servers": {
                    "srv_http": {
                        "listen": [":%s"],
                        "routes": [
                            {
                                "match": [{"host": ["%s"]}],
                                "handle": [
                                    {
                                        "handler": "reverse_proxy",
                                        "transport": {
                                            "protocol": "http",
                                            "read_timeout": 360000000000
                                        },
                                        "upstreams": [
                                            {
                                                "dial": "%s:%s"
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "handle": [
                                    {
                                        "handler": "file_server",
                                        "root": "%s",
                                        "browse": true
                                    }
                                ]
                            }
                        ],
                        "logs": {
                            "logger_names": {
                                "%s": ["log0"]
                            }
                        }
                    },
                    "srv_https": {
                        "listen": [":%s"],
                        "routes": [
                            {
                                "match": [{"host": ["%s"]}],
                                "handle": [
                                    {
                                        "handler": "reverse_proxy",
                                        "transport": {
                                            "protocol": "http",
                                            "read_timeout": 360000000000
                                        },
                                        "upstreams": [
                                            {
                                                "dial": "%s:%s"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ],
                        "logs": {
                            "logger_names": {
                                "%s": ["log1"]
                            }
                        }
                    }
                }
            },
            "tls": {
                "automation": {
                    "policies": [
                        {
                            "issuers": [
                                {
                                    "module": "internal"
                                }
                            ]
                        }
                    ]
                }
            }
        }
    }`, caddyDebugLogPath, caddyHTTPSDebugLogPath, storagePath, httpPort, httpPort, ipAddress, pbUrl, pbPort, storagePath, ipAddress, httpsPort, ipAddress, pbUrl, pbPort, ipAddress)
}
