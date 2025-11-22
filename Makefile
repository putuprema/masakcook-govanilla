.PHONY: help dev build

# Go binary paths
GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
TEMPL := $(GOBIN)/templ
AIR := $(GOBIN)/air

# Default target
help:
	@echo "Available commands:"
	@echo "  make dev           - Run development server with hot reload"
	@echo "  make build					- Build for production release"

dev:
	@echo "Starting development server with hot reload..."
	@$(AIR)

build:
	@echo "Bundling static assets..."
	@cd ./frontend && bun run build
	@echo "Building server binary..."
	@go build -o ./bin/masakcook ./cmd/server/main.go
