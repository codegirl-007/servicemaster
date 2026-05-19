ifneq (,$(wildcard .env))
include .env
export
endif

MIGRATIONS_DIR := db/migrations
DB_DRIVER := postgres

.PHONY: help run test sqlc migrate-status migrate-up migrate-down migrate-reset migrate-create

help:
	@echo "Available targets:"
	@echo "  make run                 Run the API server"
	@echo "  make test                Run Go tests"
	@echo "  make sqlc                Generate sqlc code"
	@echo "  make migrate-status      Show migration status"
	@echo "  make migrate-up          Apply all pending migrations"
	@echo "  make migrate-down        Roll back one migration"
	@echo "  make migrate-reset       Roll back all migrations"
	@echo "  make migrate-create name=<name>  Create a new SQL migration"

run:
	go run ./cmd/api

test:
	go test ./...

sqlc:
	sqlc generate

migrate-status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DATABASE_URL)" status

migrate-up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DATABASE_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DATABASE_URL)" down

migrate-reset:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DATABASE_URL)" reset

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "usage: make migrate-create name=<migration_name>"; \
		exit 1; \
	fi
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
