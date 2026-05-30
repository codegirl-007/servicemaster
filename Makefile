ifneq (,$(wildcard .env))
include .env
export
endif

MIGRATIONS_DIR := db/migrations
DB_DRIVER := postgres

.PHONY: help run test vet lint sqlc migrate-status migrate-up migrate-down migrate-reset migrate-create

help:
	@echo "Available targets:"
	@echo "  make run                 Run the API server"
	@echo "  make test                Run Go tests"
	@echo "  make vet                 Run go vet"
	@echo "  make lint                Run golangci-lint"
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

vet:
	go vet ./...

lint:
	golangci-lint run ./...

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

db-up:
	@if docker ps -a --format '{{.Names}}' | grep -q '^servicemaster-postgres$$'; then \
		echo "Starting existing Postgres container..."; \
		docker start servicemaster-postgres; \
	else \
		echo "Creating Postgres container..."; \
		docker run --name servicemaster-postgres \
			-e POSTGRES_USER=servicemaster \
			-e POSTGRES_PASSWORD=servicemaster \
			-e POSTGRES_DB=servicemaster_dev \
			-p 5432:5432 \
			-d postgres:16; \
	fi
