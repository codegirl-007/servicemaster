// Package db opens and validates the shared Postgres connection.
package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"servicemaster/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const startupPingTimeout = 5 * time.Second

// Open creates a Postgres database handle from application config.
func Open(cfg config.Config) (*sql.DB, error) {
	databaseHandle, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	configurePool(databaseHandle, cfg)

	startupContext, cancel := context.WithTimeout(context.Background(), startupPingTimeout)
	defer cancel()

	if err := Ping(startupContext, databaseHandle); err != nil {
		closeErr := databaseHandle.Close()
		if closeErr != nil {
			return nil, fmt.Errorf("ping postgres connection: %w (close failed: %v)", err, closeErr)
		}

		return nil, fmt.Errorf("ping postgres connection: %w", err)
	}

	return databaseHandle, nil
}

// Ping checks whether the database is reachable.
func Ping(ctx context.Context, databaseHandle *sql.DB) error {
	if err := databaseHandle.PingContext(ctx); err != nil {
		return fmt.Errorf("ping postgres with context: %w", err)
	}

	return nil
}

// configurePool applies connection pool settings from application config.
func configurePool(databaseHandle *sql.DB, cfg config.Config) {
	databaseHandle.SetMaxOpenConns(cfg.DatabaseMaxOpenConns)
	databaseHandle.SetMaxIdleConns(cfg.DatabaseMaxIdleConns)
	databaseHandle.SetConnMaxLifetime(cfg.DatabaseConnMaxLifetime)
	databaseHandle.SetConnMaxIdleTime(cfg.DatabaseConnMaxIdleTime)
}
