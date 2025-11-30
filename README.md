# 🚀 Golang Project Template

<div align="center">
	<strong>EN</strong> | <a href="README.ru.md">RU</a>
</div>

<br />

<div align="center">
    <a href="https://github.com/shanth1/template/actions"><img src="https://github.com/shanth1/template/actions/workflows/ci.yml/badge.svg" alt="CI Status"></a>
    <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square" alt="License"></a>
    <img src="https://img.shields.io/badge/go-1.23+-00ADD8?style=flat-square&logo=go" alt="Go Version">
</div>

<br />

This is a production-ready template for Go microservices following the **Hexagonal Architecture** (Ports and Adapters). It includes setup for configuration, logging, graceful shutdown, Docker, and CI/CD.

## 📋 Table of Contents

- [0. Prerequisites & Tools](#0-prerequisites--tools)
- [1. Getting Started](#1-getting-started)
- [2. Project Structure](#2-project-structure)
- [3. Development & Testing](#3-development--testing)
  - [Live Reload](#live-reload)
  - [Mocks Generation](#mocks-generation)
  - [Linting](#linting)
- [4. Deployment](#4-deployment)
- [5. License](#5-license)

---

## 0. Prerequisites & Tools

Before you start, ensure you have **Go (1.23+)** and **Docker** installed.

This project relies on standard Go tools for a smooth development experience. While optional, they are highly recommended.

### 🛠 Install Development Tools

**1. Air (Live Reload)**
Automatically recompiles and restarts the server when files change.

```bash
# Via Homebrew (macOS)
brew install air

# Or via Go install
go install github.com/air-verse/air@latest
```

**2. GolangCI-Lint (Linter)**
A fast Go linters runner.

```bash
# Via Homebrew (macOS)
brew install golangci-lint

# Or via Go install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
```

**3. Mockgen (Testing)**
Generates mock implementations for interfaces (essential for unit testing).

```bash
go install go.uber.org/mock/mockgen@latest
```

---

## 1. Getting Started

### 1.1. Environment Setup

Create a local `.env` file from the example:

```bash
make env
```

### 1.2. Running Locally

To run the application in development mode:

**Option A: With Live Reload (Recommended)**
Requires `air` installed.

```bash
make run-watch
```

**Option B: Standard Go Run**

```bash
make run-local
```

The server will start at `http://localhost:8080`.
Check health: `curl http://localhost:8080/health`

---

## 2. Project Structure

The project follows the standard Go project layout and Hexagonal Architecture principles.

```
├── cmd/
│   └── api/                # Application entry point (main.go)
├── config/                 # Configuration files (development, production, local)
├── deployments/            # Dockerfiles and deployment scripts
├── internal/
│   ├── adapter/            # Adapters Layer (Interaction with outside world)
│   │   ├── inbound/        # Incoming requests (HTTP REST, gRPC)
│   │   └── outbound/       # Outgoing calls (Database, External APIs)
│   ├── app/                # Application bootstrap logic
│   ├── core/               # Business Logic Layer
│   │   ├── domain/         # Domain Entities
│   │   ├── port/           # Interfaces (Ports) for Inbound/Outbound
│   │   └── service/        # Business logic implementation
│   └── lib/                # Shared internal libraries (response helpers, etc.)
└── Makefile                # Useful automation commands
```

---

## 3. Development & Testing

### Live Reload

Use `air` to watch for file changes. Configuration is located in `.air.toml`.

```bash
make run-watch
```

### Mocks Generation

We use `go generate` and `mockgen` to create mocks for testing.
Mock definitions are located inside `internal/core/port/mocks`.

To regenerate all mocks after changing interfaces:

```bash
make mocks
```

### Linting

To check code quality before committing:

```bash
make lint
```

### Running Tests

Run all unit tests with race detection:

```bash
make test
```

---

## 4. Deployment

### Docker Build

The `Dockerfile` is a multi-stage build ensuring a small and secure image (`scratch` based).

```bash
make docker-build
```

---

## 5. License

Distributed under the MIT License. See `LICENSE` for more information.
