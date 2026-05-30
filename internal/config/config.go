// Package config loads application configuration from environment variables.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config stores application configuration loaded from environment variables.
type Config struct {
	// AppEnv is the current application environment name.
	AppEnv string
	// HTTPAddr is the HTTP listen address.
	HTTPAddr string
	// DatabaseURL is the Postgres connection string.
	DatabaseURL string
	// DatabaseMaxOpenConns is the max number of open DB connections.
	DatabaseMaxOpenConns int
	// DatabaseMaxIdleConns is the max number of idle DB connections.
	DatabaseMaxIdleConns int
	// DatabaseConnMaxLifetime is the max lifetime for a DB connection.
	DatabaseConnMaxLifetime time.Duration
	// DatabaseConnMaxIdleTime is the max idle time for a DB connection.
	DatabaseConnMaxIdleTime time.Duration
	// TokenEncryptionKeyBase64 is the base64-encoded key used to encrypt stored tokens.
	TokenEncryptionKeyBase64 string
	// GoogleMapsAPIKey is to use Google Maps API
	GoogleMapsAPIKey string
}

const (
	defaultAppEnv                  = "development"
	defaultHTTPAddr                = ":8080"
	defaultDatabaseMaxOpenConns    = 10
	defaultDatabaseMaxIdleConns    = 5
	defaultDatabaseConnMaxLifetime = 30 * time.Minute
	defaultDatabaseConnMaxIdleTime = 5 * time.Minute
)

// Load reads configuration from environment variables.
func Load() (Config, error) {
	_ = godotenv.Load()

	config := Config{
		AppEnv:                   getEnv("APP_ENV", defaultAppEnv),
		HTTPAddr:                 getEnv("HTTP_ADDR", defaultHTTPAddr),
		DatabaseURL:              os.Getenv("DATABASE_URL"),
		DatabaseMaxOpenConns:     getEnvInt("DATABASE_MAX_OPEN_CONNS", defaultDatabaseMaxOpenConns),
		DatabaseMaxIdleConns:     getEnvInt("DATABASE_MAX_IDLE_CONNS", defaultDatabaseMaxIdleConns),
		DatabaseConnMaxLifetime:  getEnvDuration("DATABASE_CONN_MAX_LIFETIME", defaultDatabaseConnMaxLifetime),
		DatabaseConnMaxIdleTime:  getEnvDuration("DATABASE_CONN_MAX_IDLE_TIME", defaultDatabaseConnMaxIdleTime),
		TokenEncryptionKeyBase64: getEnv("TOKEN_ENCRYPTION_KEY_BASE64", ""),
		GoogleMapsAPIKey: getEnv("GOOGLE_MAPS_API_KEY",""),
	}

	if err := config.Validate(); err != nil {
		return Config{}, err
	}

	return config, nil
}

// Validate checks whether required configuration values are present.
func (c Config) Validate() error {
	if c.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}

	if c.TokenEncryptionKeyBase64 == "" {
		return fmt.Errorf("TOKEN_ENCRYPTION_KEY_BASE64 is required")
	}

	return nil
}

// getEnv returns an environment variable value or a default.
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

// getEnvInt returns an integer environment variable value or a default.
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}

// getEnvDuration returns a duration environment variable value or a default.
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}
