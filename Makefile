# Load .env file
include .env
export $(shell sed 's/=.*//' .env)

MIGRATE_CMD=migrate -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -path migration

.PHONY: migrate-up migrate-down

# Run Migrations
migrate-up:
	godotenv -f .env $(MIGRATE_CMD) up

# Rollback Migrations
migrate-down:
	godotenv -f .env $(MIGRATE_CMD) down 1
