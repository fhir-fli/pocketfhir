package pocketfhir

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	// Set CLI-like arguments for PocketBase to specify server address and port
	log.Printf("[DEBUG] Setting CLI arguments for server address and port: %s:%s\n", hostname, port)
	os.Args = append(os.Args[:1], "serve", "--http", fmt.Sprintf("%s:%s", hostname, port))

	// Create a configuration object with custom settings
	log.Println("[DEBUG] Creating PocketBase configuration object...")
	config := pocketbase.Config{
		DefaultDataDir:  dataDir,
		DefaultDev:      getApiLogs, // Enabling dev mode for detailed logging
		HideStartBanner: false,
	}

	// Initialize PocketBase with the default configuration
	log.Println("[DEBUG] Initializing PocketBase app...")
	app := pocketbase.NewWithConfig(config)
	log.Println("[DEBUG] PocketBase app initialized.")

	// Standard setup for the app
	log.Println("[DEBUG] Running standard setup for PocketBase...")
	standard(app)
	log.Println("[DEBUG] Standard setup for PocketBase app completed.")

	// Register hooks from hooks.go
	log.Println("[DEBUG] Registering hooks...")
	registerHooks(app)
	log.Println("[DEBUG] Hooks registered.")

	// Register FHIR routes
	log.Println("[DEBUG] Registering FHIR routes...")
	registerFHIRRoutes(app)
	log.Println("[DEBUG] FHIR routes registered.")

	// Register server management routes
	log.Println("[DEBUG] Registering server management routes...")
	registerManagementRoutes(app)
	log.Println("[DEBUG] Server management routes registered.")

	// Setup additional native callbacks and routes
	log.Println("[DEBUG] Setting up PocketBase callbacks and routes...")
	setupPocketbaseCallbacks(app, getApiLogs)
	log.Println("[DEBUG] PocketBase callbacks and routes set up.")

	// Initialize collections if necessary
	log.Println("[DEBUG] Initializing collections...")
	if err := initializeCollections(app); err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	} else {
		log.Println("[DEBUG] Collections initialized successfully.")
	}

	// Start the server (no need to wait on anything here)
	log.Println("[DEBUG] Calling app.Start() to start the server...")
	if err := app.Start(); err != nil {
		sendCommand("error", fmt.Sprintf("Error: Failed to start PocketBase server: %v", err))
		log.Fatalf("Failed to start the app: %v", err)
	} else {
		log.Println("[DEBUG] PocketFHIR server started successfully.")
	}
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
		log.Printf("[DEBUG] Sending command '%s' with data: %s", command, data)
		return nativeBridge.HandleCallback(command, data)
	}
	log.Printf("[DEBUG] No NativeBridge defined. Command '%s' not sent.", command)
	return ""
}

// setupPocketbaseCallbacks sets up additional callbacks and native routes for PocketBase
func setupPocketbaseCallbacks(app *pocketbase.PocketBase, getApiLogs bool) {
	// Setup callbacks
	log.Println("[DEBUG] Setting up OnBeforeServe callback...")
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		log.Println("[DEBUG] OnBeforeServe triggered.")
		sendCommand("OnBeforeServe", "")
		if getApiLogs {
			log.Println("[DEBUG] Adding API logs middleware...")
			e.Router.Use(ApiLogsMiddleWare(app))
		}

		// Setup a native GET request handler
		log.Println("[DEBUG] Adding native GET request handler...")
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/nativeGet",
			Handler: func(context echo.Context) error {
				log.Println("[DEBUG] Handling native GET request...")
				var data = sendCommand("nativeGetRequest", context.QueryParams().Encode())
				return context.JSON(http.StatusOK, map[string]string{
					"success": data,
				})
			},
		})

		// Setup a native POST request handler
		log.Println("[DEBUG] Adding native POST request handler...")
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/nativePost",
			Handler: func(context echo.Context) error {
				log.Println("[DEBUG] Handling native POST request...")
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
		log.Println("[DEBUG] OnBeforeBootstrap triggered.")
		sendCommand("OnBeforeBootstrap", "")
		return nil
	})
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		log.Println("[DEBUG] OnAfterBootstrap triggered.")
		sendCommand("OnAfterBootstrap", "")
		return nil
	})
	app.OnTerminate().Add(func(e *core.TerminateEvent) error {
		log.Println("[DEBUG] OnTerminate triggered.")
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
			log.Printf("[DEBUG] API request made to: %s", fullPath)
			sendCommand("apiLogs", fullPath)
			return next(c)
		}
	}
}
