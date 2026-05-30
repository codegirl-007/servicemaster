package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"servicemaster/internal/config"
	"servicemaster/internal/platform/db"
	"syscall"
	"time"
)

func main() {

	// load configs
	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}

	slog.Info("initializing database")

	databaseHandle, err := db.Open(cfg)
	if err != nil {
		slog.Error("initializing database", "error", err)
		os.Exit(1)
	}
	defer databaseHandle.Close()

	slog.Info("database initialized")

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := db.Ping(ctx, databaseHandle); err != nil {
			http.Error(w, "database unavailable", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready\n"))
	})

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		slog.Info("starting http server", "addr", cfg.HTTPAddr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdownSignals)

	select {
	case err := <-serverErrors:
		slog.Error("http server", "error", err)
		os.Exit(1)
	case s := <-shutdownSignals:
		slog.Info("received shutdown signal", "signal", s)
	}

	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		slog.Error("shutdown http server", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")

}
