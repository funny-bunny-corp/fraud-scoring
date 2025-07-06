# Fraud Scoring System Makefile

# Variables
BINARY_NAME=fraud-scoring
MAIN_PATH=cmd/main.go
BUILD_DIR=bin
PROTO_DIR=api
COVERAGE_DIR=coverage

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint
WIRE=wire

# Default target
.PHONY: all
all: clean deps generate build test

# Build the application
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -rf $(COVERAGE_DIR)

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Generate code (protobuf, wire)
.PHONY: generate
generate:
	@echo "Generating code..."
	@protoc --go_out=. --go-grpc_out=. $(PROTO_DIR)/payment-processing.proto
	@cd cmd && $(WIRE)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p $(COVERAGE_DIR)
	$(GOTEST) -coverprofile=$(COVERAGE_DIR)/coverage.out ./...
	$(GOCMD) tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@echo "Coverage report generated at $(COVERAGE_DIR)/coverage.html"

# Run tests with race detection
.PHONY: test-race
test-race:
	@echo "Running tests with race detection..."
	$(GOTEST) -race ./...

# Run benchmarks
.PHONY: bench
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Lint code
.PHONY: lint
lint:
	@echo "Running linter..."
	$(GOLINT) run

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	$(GOCMD) vet ./...

# Security check
.PHONY: security
security:
	@echo "Running security check..."
	@command -v gosec >/dev/null 2>&1 || { echo "gosec not found, installing..."; go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest; }
	gosec ./...

# Run the application
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Build Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME) .

# Run Docker container
.PHONY: docker-run
docker-run: docker-build
	@echo "Running Docker container..."
	docker run --rm -p 50051:50051 $(BINARY_NAME)

# Install development tools
.PHONY: install-tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# Pre-commit checks
.PHONY: pre-commit
pre-commit: fmt vet lint security test

# Development workflow
.PHONY: dev
dev: deps generate pre-commit build

# Full CI pipeline
.PHONY: ci
ci: clean deps generate fmt vet lint security test-coverage test-race build

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all          - Clean, deps, generate, build, test"
	@echo "  build        - Build the application"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Download dependencies"
	@echo "  generate     - Generate code (protobuf, wire)"
	@echo "  test         - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  test-race    - Run tests with race detection"
	@echo "  bench        - Run benchmarks"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  security     - Run security check"
	@echo "  run          - Run the application"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  install-tools - Install development tools"
	@echo "  pre-commit   - Run pre-commit checks"
	@echo "  dev          - Development workflow"
	@echo "  ci           - Full CI pipeline"
	@echo "  help         - Show this help"