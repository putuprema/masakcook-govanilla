.PHONY: help dev

# Go binary paths
GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
TEMPL := $(GOBIN)/templ
AIR := $(GOBIN)/air

# Default target
help:
	@echo "Available commands:"
	@echo "  make dev           - Run development server with hot reload"

dev:
	@echo "Starting development server with hot reload..."
	@$(AIR)
