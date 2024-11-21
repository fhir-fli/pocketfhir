package pocketfhir

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

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
func RunServer(dataDir string, ipAddress string, pbPort string, enableApiLogs bool) {
	// Set CLI-like arguments for PocketBase to specify server address and port
	log.Printf("[DEBUG] Setting CLI arguments for server address and port: %s:%s\n", ipAddress, pbPort)
	os.Args = []string{os.Args[0], "serve", "--http", fmt.Sprintf("%s:%s", ipAddress, pbPort)}

	// Create a configuration object with custom settings
	log.Println("[DEBUG] Creating PocketBase configuration object...")
	config := pocketbase.Config{
		DefaultDataDir:  dataDir,
		DefaultDev:      enableApiLogs, // Enabling dev mode for detailed logging
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
	setupPocketbaseCallbacks(app, enableApiLogs)
	log.Println("[DEBUG] PocketBase callbacks and routes set up.")

	// Initialize collections if necessary
	log.Println("[DEBUG] Initializing collections...")
	if err := initializeCollections(app); err != nil {
		log.Printf("[ERROR] Failed to initialize collections: %v", err)
		return
	}
	log.Println("[DEBUG] Collections initialized successfully.")

	// Start the server in a separate goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("[DEBUG] Calling app.Start() to start the server...")
		if err := app.Start(); err != nil {
			sendCommand("error", fmt.Sprintf("Error: Failed to start PocketBase server: %v", err))
			log.Printf("[ERROR] Failed to start the app: %v", err)
		} else {
			log.Println("[DEBUG] PocketFHIR server started successfully.")
		}
	}()

	// Wait for server to complete
	wg.Wait()
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
func setupPocketbaseCallbacks(app *pocketbase.PocketBase, enableApiLogs bool) {
	setupOnBeforeServe(app, enableApiLogs)
	setupOnBeforeBootstrap(app)
	setupOnAfterBootstrap(app)
	setupOnTerminate(app)
}

func setupOnBeforeServe(app *pocketbase.PocketBase, enableApiLogs bool) {
	log.Println("[DEBUG] Setting up OnBeforeServe callback...")
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		log.Println("[DEBUG] OnBeforeServe triggered.")
		sendCommand("OnBeforeServe", "")
		if enableApiLogs {
			log.Println("[DEBUG] Adding API logs middleware...")
			e.Router.Use(ApiLogsMiddleWare(app))
		}

		setupNativeGetHandler(e.Router)
		setupNativePostHandler(e.Router)
		return nil
	})
}

func setupNativeGetHandler(router *echo.Echo) {
	log.Println("[DEBUG] Adding native GET request handler...")
	router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/api/nativeGet",
		Handler: func(context echo.Context) error {
			log.Println("[DEBUG] Handling native GET request...")
			data := sendCommand("nativeGetRequest", context.QueryParams().Encode())
			return context.JSON(http.StatusOK, map[string]string{"success": data})
		},
	})
}

func setupNativePostHandler(router *echo.Echo) {
	log.Println("[DEBUG] Adding native POST request handler...")
	router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/api/nativePost",
		Handler: func(context echo.Context) error {
			log.Println("[DEBUG] Handling native POST request...")
			form, err := context.FormValues()
			if err != nil {
				return context.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}
			data := sendCommand("nativePostRequest", form.Encode())
			return context.JSON(http.StatusOK, map[string]string{"success": data})
		},
	})
}

func setupOnBeforeBootstrap(app *pocketbase.PocketBase) {
	log.Println("[DEBUG] Setting up OnBeforeBootstrap callback...")
	app.OnBeforeBootstrap().Add(func(e *core.BootstrapEvent) error {
		log.Println("[DEBUG] OnBeforeBootstrap triggered.")
		sendCommand("OnBeforeBootstrap", "")
		return nil
	})
}

func setupOnAfterBootstrap(app *pocketbase.PocketBase) {
	log.Println("[DEBUG] Setting up OnAfterBootstrap callback...")
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		log.Println("[DEBUG] OnAfterBootstrap triggered.")
		sendCommand("OnAfterBootstrap", "")
		return nil
	})
}

func setupOnTerminate(app *pocketbase.PocketBase) {
	log.Println("[DEBUG] Setting up OnTerminate callback...")
	app.OnTerminate().Add(func(e *core.TerminateEvent) error {
		log.Println("[DEBUG] OnTerminate triggered.")
		sendCommand("OnTerminate", "")
		return nil
	})
}

// ApiLogsMiddleWare logs all API calls for PocketBase
var requestLogBuffer []string
var mu sync.Mutex

func init() {
	// Background goroutine to batch log requests
	go func() {
		for {
			time.Sleep(1 * time.Second)
			mu.Lock()
			if len(requestLogBuffer) > 0 {
				log.Println("[DEBUG] Batch API request logs:", requestLogBuffer)
				requestLogBuffer = []string{} // Clear the buffer
			}
			mu.Unlock()
		}
	}()
}

func ApiLogsMiddleWare(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			fullPath := request.URL.Host + request.URL.Path + "?" + request.URL.RawQuery
			mu.Lock()
			requestLogBuffer = append(requestLogBuffer, fullPath)
			mu.Unlock()
			return next(c)
		}
	}
}
