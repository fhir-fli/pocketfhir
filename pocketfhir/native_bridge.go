package pocketfhir

import (
	"log"
)

// Java Callbacks, make sure to register them before starting pocketbase
// to expose any method to java, add that with FirstLetterCapital
var nativeBridge NativeBridge

// RegisterNativeBridgeCallback allows setting the NativeBridge interface for callbacks
func RegisterNativeBridgeCallback(c NativeBridge) {
	nativeBridge = c
}

// Helper methods

// NativeBridge interface to define the methods used for native callbacks
type NativeBridge interface {
	handleCallback(string, string) string
}

// sendCommand sends command to native and returns the response
func sendCommand(command string, data string) string {
	if nativeBridge != nil {
		log.Printf("[DEBUG] Sending command '%s' with data: %s", command, data)
		return nativeBridge.handleCallback(command, data)
	}
	log.Printf("[DEBUG] No NativeBridge defined. Command '%s' not sent.", command)
	return ""
}
