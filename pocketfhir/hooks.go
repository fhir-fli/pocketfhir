package pocketfhir

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// Register custom hooks
func registerHooks(app *pocketbase.PocketBase) {
	// Register middleware to log each request
	registerRequestLoggingHook(app)

	// Setup versioning and history hooks
	setupHooks(app)
}

// Log each request
func registerRequestLoggingHook(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				log.Printf("Request: %s %s", c.Request().Method, c.Request().URL.Path)
				return next(c)
			}
		})
		return nil
	})
}

func setupHooks(app *pocketbase.PocketBase) {
	// Hook for handling resource creation
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		return handleResourceCreation(app, e)
	})

	// Hook for handling resource updates
	app.OnRecordBeforeUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		return handleResourceUpdate(app, e)
	})
}
