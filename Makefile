include .env
export

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR=migrations

.PHONY: run build run-build test migration migrate-up

run:
	go run ./cmd/helpdesk

build:
	go build -o ./bin/helpdesk ./cmd/helpdesk

run-build:
	./bin/helpdesk

test:
	go test -v ./...

migration:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migration name=create_users_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

