package config

import (
	"os"
)

// PORT is the default port the server listens on.
const PORT = ":8081"

// GetPort retrieves the port from environment variable PORT or falls back to the default.
func GetPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return PORT
}
