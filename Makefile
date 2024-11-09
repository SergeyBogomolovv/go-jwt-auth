include .env
MIGRATIONS_PATH = ./cmd/migrations

.PHONY: build
build:
	@go build -o bin/auth cmd/main.go

.PHONY: test
test:
	@go test -v ./...

.PHONY: run
run: build
	@./bin/auth

.PHONY:	migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URL) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URL) down $(filter-out $@,$(MAKECMDGOALS))