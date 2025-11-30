.PHONY: help test lint clean build run docker-build

# --- Project Variables ---
BINARY_NAME := goproject
CMD_API_PATH := ./cmd/api
CONFIG_FILE := config/local.yaml

# --- Build Variables ---
COMMIT_HASH := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME := $(shell date +%FT%T%z)
LDFLAGS := -ldflags "-X main.CommitHash=$(COMMIT_HASH) -X main.BuildTime=$(BUILD_TIME) -s -w"

# --- Docker Variables ---
DOCKER_TAG := $(COMMIT_HASH)

help: ## Show this help message
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1mUsage:\033[0m\n  make \033[36m<target>\033[0m\n"} \
	/^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } \
	/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Environment

env: ## Create .env file from example
	@cp -n .env.example .env || true

##@ Development

run-local: ## Run app in local mode
	@go run $(CMD_API_PATH) --config $(CONFIG_FILE)

run-watch: ## Run with live reload (requires 'air' in PATH)
	@air -c .air.toml

mocks: ## Generate all mocks
	@echo "Generating mocks..."
	@go generate ./...

##@ Testing & Quality

test: ## Run unit tests
	@go test -v -race ./...

lint: ## Run golangci-lint
	golangci-lint run ./...

lint-install: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

##@ Builds

build: ## Build binary for current OS
	@echo "Building $(BINARY_NAME)..."
	@go build $(LDFLAGS) -o build/$(BINARY_NAME) $(CMD_API_PATH)
	@echo "Build complete: build/$(BINARY_NAME)"

build-linux: ## Build binary for Linux (AMD64)
	@echo "Building Linux binary..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o build/$(BINARY_NAME)-linux $(CMD_API_PATH)
	@echo "Build complete: build/$(BINARY_NAME)-linux"

clean: ## Remove build artifacts
	@rm -rf build/

##@ Deployment

docker-build: ## Build docker image with commit tag
	@echo "Building Docker image $(BINARY_NAME):$(DOCKER_TAG)..."
	@docker build -t $(BINARY_NAME):$(DOCKER_TAG) -t $(BINARY_NAME):latest -f deployments/Dockerfile .
