package logger

import (
	"log"
	"os"
)

// Global loggers for the entire application
var (
	Info  *log.Logger // For informational messages
	Error *log.Logger // For error messages
	Debug *log.Logger // For debug messages
)

// Init initializes the logging system
func Init(level string) error {
	// Basic formatting settings
	flags := log.Ldate | log.Ltime | log.Lshortfile

	// Initialize loggers
	Info = log.New(os.Stdout, "INFO: ", flags)
	Error = log.New(os.Stderr, "ERROR: ", flags)
	Debug = log.New(os.Stdout, "DEBUG: ", flags)

	// TODO: add support for different log levels
	log.Printf("Logger initialized with level: %s", level)

	return nil
}
