# Go parameters
APP_NAME=kairosai-backend
BUILD_DIR=bin
GO_FILES=$(shell find . -name '*.go' -type f)

# Environment variables
ENV_FILE=.env
PORT=$(shell grep PORT $(ENV_FILE) | cut -d '=' -f2)
DB_HOST=$(shell grep DB_HOST $(ENV_FILE) | cut -d '=' -f2)
DB_PORT=$(shell grep DB_PORT $(ENV_FILE) | cut -d '=' -f2)
DB_USER=$(shell grep DB_USER $(ENV_FILE) | cut -d '=' -f2)
DB_PASSWORD=$(shell grep DB_PASSWORD $(ENV_FILE) | cut -d '=' -f2)
DB_NAME=$(shell grep DB_NAME $(ENV_FILE) | cut -d '=' -f2)

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make run         - Run the Go server"
	@echo "  make build       - Build the Go binary"
	@echo "  make clean       - Remove generated files"
	@echo "  make test        - Run unit tests"

# Run the Go server
.PHONY: run
run:
	@echo "Starting KairosAI Backend Service..."
	@go run main.go

# Build the Go binary
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "Build complete! Binary is at $(BUILD_DIR)/$(APP_NAME)"

# Remove built files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned!"

# Run unit tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -cover -race
