package main

import (
	"context"
	"errors"
	"log"
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
		log.Fatalf("load config: %v", err)
	}

	log.Println("initializing database...")

	databaseHandle, err := db.Open(cfg)
	if err != nil {
		log.Fatalf("initializing database: %v", err)
	}
	defer databaseHandle.Close()

	log.Println("database initialized")

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
		Addr: cfg.HTTPAddr,
		Handler: mux,
		ReadHeaderTimeout: time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("starting http server on %s", cfg.HTTPAddr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdownSignals)

	select {
	case err := <-serverErrors:
		log.Fatalf("http server: %v", err)
	case signal := <- shutdownSignals:
		log.Printf("received shutdown signal: %s", signal)
	}

	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		log.Fatalf("shutdown http server: %v", err)
	}

	log.Println("server stopped")

}
