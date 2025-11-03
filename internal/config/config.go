package config

import (
	"os"
	"strconv"
)

// Config contains all application configuration
type Config struct {
	Server   ServerConfig   // Server settings
	Database DatabaseConfig // Database settings
	JWT      JWTConfig      // JWT settings
	Log      LogConfig      // Logging settings
	Auth     AuthConfig     // Authentication settings
}

// ServerConfig contains HTTP server configuration
type ServerConfig struct {
	Port string // Server port
}

// DatabaseConfig contains database connection settings
type DatabaseConfig struct {
	URL string // Database connection URL
}

// JWTConfig contains JWT token settings
type JWTConfig struct {
	Secret string // JWT signing secret
}

// LogConfig contains logging configuration
type LogConfig struct {
	Level  string // Log level (debug, info, warn, error)
	Format string // Log format (text, json)
}

// AuthConfig contains authentication settings
type AuthConfig struct {
	BcryptCost int // Password hashing cost
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{}

	// Server settings
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// Database settings
	cfg.Database.URL = getEnv("DATABASE_URL", "postgres://tabloom_user:tabloom_password@localhost:5432/tabloom?sslmode=disable")

	// JWT settings
	cfg.JWT.Secret = getEnv("JWT_SECRET", "tabloom-default-secret-key-change-in-production")

	// Logging settings
	cfg.Log.Level = getEnv("LOG_LEVEL", "info")
	cfg.Log.Format = getEnv("LOG_FORMAT", "text")

	// Authentication settings
	cfg.Auth.BcryptCost = getEnvInt("BCRYPT_COST", 12)

	return cfg, nil
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt gets environment variable as integer or returns default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
