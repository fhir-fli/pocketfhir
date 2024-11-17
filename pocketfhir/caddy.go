package pocketfhir

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caddyserver/caddy/v2"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

func CreateConfig(port, certFile, keyFile, upstreamURL, logFile, adminListen string) *caddy.Config {
	return &caddy.Config{
		Admin: &caddy.AdminConfig{
			Listen: adminListen,
		},
		AppsRaw: caddy.ModuleMap{
			"http": json.RawMessage(fmt.Sprintf(`{
                "servers": {
                    "reverse_proxy_server": {
                        "listen": [":%s"],
                        "tls_connection_policies": [
                            {
                                "certificates": {
                                    "load_files": [
                                        {
                                            "certificate_file": "%s",
                                            "key_file": "%s"
                                        }
                                    ]
                                }
                            }
                        ],
                        "routes": [
                            {
                                "match": [{"host": ["*"]}],
                                "handle": [{"handler": "reverse_proxy", "upstreams": [{"dial": "%s"}]}]
                            },
                            {
                                "match": [{"path": ["/health"]}],
                                "handle": [{"handler": "static_response", "body": "PocketFHIR Proxy Active", "status_code": 200}]
                            }
                        ],
                        "logs": {
                            "default_logger_name": "access_log"
                        }
                    }
                }
            }`, port, certFile, keyFile, upstreamURL)),
		},
		Logging: &caddy.Logging{
			Logs: map[string]*caddy.CustomLog{
				"access_log": {
					BaseLog: caddy.BaseLog{
						WriterRaw: json.RawMessage(fmt.Sprintf(`{
                            "output": {
                                "module": "file", // Add this line
                                "format": "file",
                                "filename": "%s"
                            }
                        }`, logFile)),
						Level: "INFO",
					},
				},
			},
		},
	}
}

func StartCaddy(port, certFile, keyFile, upstreamURL, logFile, adminListen, caddyConfig string) {
	// Use the provided dataPath to set the config directory
	if err := os.MkdirAll(caddyConfig, 0755); err != nil {
		log.Fatalf("Failed to create configuration directory at %s: %v", caddyConfig, err)
	}

	log.Printf("Starting Caddy with configuration:\n"+
		"Port: %s\nCertFile: %s\nKeyFile: %s\nUpstreamURL: %s\nLogFile: %s\nAdminListen: %s\nConfigDir: %s\n",
		port, certFile, keyFile, upstreamURL, logFile, adminListen, caddyConfig)

	// Set the Caddy config directory
	errCaddyConfig := os.Setenv("CADDY_CONFIG_DIR", caddyConfig)
	if errCaddyConfig != nil {
		log.Fatalf("Failed to set CADDY_CONFIG_DIR environment variable: %v", errCaddyConfig)
	}
	errCaddyHome := os.Setenv("XDG_CONFIG_HOME", caddyConfig)
	if errCaddyHome != nil {
		log.Fatalf("Failed to set XDG_CONFIG_HOME environment variable: %v", errCaddyHome)
	}

	cfg := CreateConfig(port, certFile, keyFile, upstreamURL, logFile, adminListen)
	configJSON, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalf("Failed to serialize Caddy config: %v", err)
	}
	log.Printf("Generated Caddy config: %s", string(configJSON))

	// Start the Caddy server
	log.Println("Initializing Caddy...")
	if err := caddy.Run(cfg); err != nil {
		log.Fatalf("Error running Caddy: %v", err)
	}

	log.Println("Caddy server started successfully.")
}
