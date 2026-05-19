package main

import (
	"log"
	"servicemaster/internal/config"
	"servicemaster/internal/platform/db"
)

func main() {

	// load configs
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	log.Println("initializing database...")

	db, err := db.Open(cfg)
	if err != nil {
		log.Fatalf("initializing database: %v", err)
	}
	defer db.Close()

	log.Println("database initialized")
}
