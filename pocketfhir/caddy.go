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
	// Define paths for certs and proxy logs
	certsPath := fmt.Sprintf("%s/certs_access.log", storagePath)
	pbProxyPath := fmt.Sprintf("%s/proxy_access.log", storagePath)
	rootCertPath := fmt.Sprintf("%s/pki/authorities/local", storagePath)

	// Generate the JSON configuration
	jsonConfig := jsonConfig(pbPort, httpPort, httpsPort, pbUrl, storagePath, ipAddress, certsPath, pbProxyPath, rootCertPath)
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

func jsonConfig(pbPort, httpPort, httpsPort, pbUrl, storagePath, ipAddress, certsPath, pbProxyPath, rootCertPath string) string {
	// Construct each part individually and combine them
	loggingConfig := generateLoggingConfig(certsPath, pbProxyPath)
	storageConfig := generateStorageConfig(storagePath)
	httpAppConfig := generateHttpAppConfig(pbPort, httpPort, httpsPort, pbUrl, rootCertPath, ipAddress)
	tlsConfig := generateTlsConfig(ipAddress)

	// Combine the parts into the final JSON configuration
	return fmt.Sprintf(`{
        %s,
        %s,
        "apps": {
            %s,
            %s
        }
    }`, loggingConfig, storageConfig, httpAppConfig, tlsConfig)
}

// Generates the logging configuration section
func generateLoggingConfig(certsPath, pbProxyPath string) string {
	return fmt.Sprintf(`"logging": {
            "logs": {
                "certs": {
                    "writer": {
                        "output": "file",
                        "filename": "%s"
                    },
                    "encoder": {
                        "format": "json"
                    },
                    "include": [
                        "http.log.access.certs"
                    ]
                },
                "pb_proxy": {
                    "writer": {
                        "output": "file",
                        "filename": "%s"
                    },
                    "encoder": {
                        "format": "json"
                    },
                    "include": [
                        "http.log.access.pb_proxy"
                    ]
                }
            }
        }`, certsPath, pbProxyPath)
}

// Generates the storage configuration section
func generateStorageConfig(storagePath string) string {
	return fmt.Sprintf(`"storage": {
            "module": "file_system",
            "root": "%s"
        }`, storagePath)
}

// Generates the HTTP application configuration section
func generateHttpAppConfig(pbPort, httpPort, httpsPort, pbUrl, rootCertPath, ipAddress string) string {
	httpServerConfig := generateHttpServerConfig(httpPort, ipAddress)
	httpsServerConfig := generateHttpsServerConfig(httpsPort, ipAddress, pbUrl, pbPort)
	return fmt.Sprintf(`"http": {
            "http_port": %s,
            "servers": {
                %s,
                %s
            }
        }`, httpPort, httpServerConfig, httpsServerConfig)
}

// Generates the HTTP server configuration without static file server handling
func generateHttpServerConfig(httpPort, ipAddress string) string {
	return fmt.Sprintf(`"srv1": {
            "listen": [":%s"],
            "routes": [
                {
                    "match": [{"host": ["%s"]}],
                    "handle": [
                        {
                            "handler": "subroute"
                        }
                    ],
                    "terminal": true
                }
            ]
        }`, httpPort, ipAddress)
}

// Generates the HTTPS server configuration
// Generates the HTTP server configuration for reverse proxy with health checks
func generateHttpsServerConfig(httpsPort, ipAddress, pbUrl, pbPort string) string {
	return fmt.Sprintf(`"srv_https": {
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
                            "health_checks": {
                                "active": {
                                    "interval": 10000000000,
                                    "timeout": 2000000000,
                                    "uri": "/api/health"
                                }
                            },
                            "upstreams": [
                                {
                                    "dial": "%s:%s"
                                }
                            ]
                        }
                    ]
                }
            ]
        }`, httpsPort, ipAddress, pbUrl, pbPort)
}

// Generates the TLS automation configuration
func generateTlsConfig(ipAddress string) string {
	return fmt.Sprintf(`"tls": {
            "automation": {
                "policies": [
                    {
                        "subjects": [
                            "%s"
                        ],
                        "issuers": [
                            {
                                "module": "internal"
                            }
                        ]
                    }
                ]
            }
        }`, ipAddress)
}
