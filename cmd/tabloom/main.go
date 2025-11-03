package main

import (
	"log"

	"github.com/Bl00mGuy/Tabloom/internal/app"
)

// Application version info (populated during build)
var (
	version = "dev"     // Application version
	commit  = "none"    // Git commit hash
	date    = "unknown" // Build date
)

func main() {
	// Log version information
	log.Printf("Starting Tabloom v%s (commit: %s, built: %s)", version, commit, date)

	// Create application instance
	application, err := app.New(version, commit)
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}

	// Run the application
	if err := application.Run(); err != nil {
		log.Fatalf("Application runtime error: %v", err)
	}
}
