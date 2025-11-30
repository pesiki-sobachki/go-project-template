.PHONY: help test lint fmt vet tidy clean build run

CMD_API_PATH := ./cmd/api
BINARY_NAME := goproject
CONFIG_FILE := config/local.yaml

# List of all Go files to track for changes
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*")

help: ## Show this help message
	@echo "Usage: make <command>"
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# --- ENV ---

env: ## Safely creates a .env file based on .env.example
	@cp -n .env.example .env || true

# --- BUILDS ---

build: $(GOFILES) ## Builds api application
	@echo "Building $(BINARY_NAME)..."
	@go build -o build/$(BINARY_NAME) $(CMD_API_PATH)
	@echo "Build of build/$(BINARY_NAME) complete."

build-linux: $(GOFILES) ## Builds api application for linux
	@echo "Building $(BINARY_NAME) for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o build/$(BINARY_NAME) $(CMD_API_PATH)
	@echo "Build of build/$(BINARY_NAME) complete."

# --- RUN ---

run-local: ## Runs applications in local development mode
	@go run cmd/api/main.go --config $(CONFIG_FILE)

AIR_BIN := $(shell go env GOPATH)/bin/air
run-watch: ## Run with live reload
	$(AIR_BIN) -c .air.toml
# --- TEST ---

test: ## Runs tests
	@go test -v ./...

# --- COMMON ---

clean: ## Remove build artifacts
	@rm -rf build/

# --- DEPLOY ---

# TODO: docker-compose

docker-build: ## Build docker image
	@docker build -t $(BINARY_NAME) -f deployments/Dockerfile .


# --- LINTER ---

lint: ## Run golangci-lint
	golangci-lint run ./...

lint-install: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
