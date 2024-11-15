package pocketfhir

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// Java Callbacks, make sure to register them before starting pocketbase
// to expose any method to java, add that with FirstLetterCapital
var nativeBridge NativeBridge

// RegisterNativeBridgeCallback allows setting the NativeBridge interface for callbacks
func RegisterNativeBridgeCallback(c NativeBridge) {
	nativeBridge = c
}

// RunServer starts the PocketFHIR server
func RunServer(dataDir string, hostname string, port string, getApiLogs bool) {
	// Set environment variables for PocketBase configuration
	if err := os.Setenv("POCKETBASE_DATA_DIR", dataDir); err != nil {
		log.Fatalf("Failed to set data directory: %v", err)
	}

	// Set CLI-like arguments for PocketBase to specify server address and port
	// This will work similar to calling the PocketBase server from CLI
	os.Args = append(os.Args[:1], "serve", "--http", fmt.Sprintf("%s:%s", hostname, port))

	// Create a configuration object with custom settings
	config := pocketbase.Config{
		DefaultDataDir:  dataDir,
		DefaultDev:      getApiLogs, // Enabling dev mode for detailed logging
		HideStartBanner: false,
	}

	// Initialize PocketBase with the default configuration
	log.Println("Initializing PocketBase app...")
	app := pocketbase.NewWithConfig(config)

	// Standard setup for the app
	standard(app)
	log.Println("Standard setup for PocketBase app completed.")

	// Register hooks from hooks.go
	registerHooks(app)
	log.Println("Hooks registered.")

	// Register FHIR routes
	registerFHIRRoutes(app)
	log.Println("FHIR routes registered.")

	// Register server management routes
	registerManagementRoutes(app)
	log.Println("Server management routes registered.")

	// Setup additional native callbacks and routes
	setupPocketbaseCallbacks(app, getApiLogs)

	// Initialize collections if necessary
	log.Println("Initializing collections...")
	if err := initializeCollections(app); err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	} else {
		log.Println("Collections initialized successfully.")
	}

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		log.Println("Starting PocketFHIR server...")
		serverUrl := fmt.Sprintf("http://%s:%s", hostname, port)
		sendCommand("onServerStarting", fmt.Sprintf("Server starting at: %s\n➜ REST API: %s/api/\n➜ Admin UI: %s/_/", serverUrl, serverUrl, serverUrl))

		// Use the Start() method to start the PocketBase server
		if err := app.Start(); err != nil {
			sendCommand("error", fmt.Sprintf("Error: Failed to start PocketBase server: %v", err))
			log.Fatalf("Failed to start the app: %v", err)
		} else {
			log.Println("PocketFHIR server started successfully.")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	<-stop
	log.Println("Shutting down PocketFHIR server...")
	app.ResetBootstrapState() // Optional: Clean up any resources before shutting down
	log.Println("PocketFHIR server shut down gracefully.")
}

// StopServer gracefully stops the running PocketFHIR server
func StopServer() {
	log.Println("Stopping PocketFHIR server...")
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

// Helper methods

// NativeBridge interface to define the methods used for native callbacks
type NativeBridge interface {
	HandleCallback(string, string) string
}

// sendCommand sends command to native and returns the response
func sendCommand(command string, data string) string {
	if nativeBridge != nil {
		return nativeBridge.HandleCallback(command, data)
	}
	return ""
}

// setupPocketbaseCallbacks sets up additional callbacks and native routes for PocketBase
func setupPocketbaseCallbacks(app *pocketbase.PocketBase, getApiLogs bool) {
	// Setup callbacks
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		sendCommand("OnBeforeServe", "")
		if getApiLogs {
			e.Router.Use(ApiLogsMiddleWare(app))
		}

		// Setup a native GET request handler
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/nativeGet",
			Handler: func(context echo.Context) error {
				var data = sendCommand("nativeGetRequest", context.QueryParams().Encode())
				return context.JSON(http.StatusOK, map[string]string{
					"success": data,
				})
			},
		})

		// Setup a native POST request handler
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/nativePost",
			Handler: func(context echo.Context) error {
				form, err := context.FormValues()
				if err != nil {
					return context.JSON(http.StatusBadRequest, map[string]string{
						"error": err.Error(),
					})
				}
				var data = sendCommand("nativePostRequest", form.Encode())
				return context.JSON(http.StatusOK, map[string]string{
					"success": data,
				})
			},
		})

		return nil
	})

	app.OnBeforeBootstrap().Add(func(e *core.BootstrapEvent) error {
		sendCommand("OnBeforeBootstrap", "")
		return nil
	})
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		sendCommand("OnAfterBootstrap", "")
		return nil
	})
	app.OnTerminate().Add(func(e *core.TerminateEvent) error {
		sendCommand("OnTerminate", "")
		return nil
	})
}

// ApiLogsMiddleWare logs all API calls for PocketBase
func ApiLogsMiddleWare(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			fullPath := request.URL.Host + request.URL.Path + "?" + request.URL.RawQuery
			sendCommand("apiLogs", fullPath)
			return next(c)
		}
	}
}
