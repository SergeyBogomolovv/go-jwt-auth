include .env

.PHONY: build
build:
	@go build -o bin/auth cmd/main.go

.PHONY: test
test:
	@go test -v ./...

.PHONY: run
run: build
	@./bin/auth