# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod

# Proto files directory and generated Go files directory
PROTO_DIR = api
PROTO_GO_OUT = api

# Protoc command and options
PROTOC = protoc
PROTOC_INCLUDES = -I$(PROTO_DIR)
PROTOC_GO_OPTIONS = --go_out=$(PROTO_GO_OUT) --go-grpc_out=$(PROTO_GO_OUT)


# Main application path
MAIN_PATH = cmd/main.go

# Name of the binary executable
BINARY_NAME = grpc-client

# Directory containing all Go source files
SRC_DIR = ./...

# Targets

# Default target
default: build

# Build the binary
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) $(MAIN_PATH)

# Clean up the binary
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)

# Install dependencies
deps:
	$(GOGET) ./...

# Run tests
test:
	$(GOTEST) $(SRC_DIR) -v

# Update and tidy dependencies
mod:
	$(GOMOD) tidy

# Run the application
run:
	$(GOBUILD) -o bin/$(BINARY_NAME) $(MAIN_PATH)
	./bin/$(BINARY_NAME)

# Generate Go code from proto files
generate-proto:
	$(PROTOC) $(PROTOC_INCLUDES) $(PROTOC_GO_OPTIONS) $(PROTO_DIR)/*.proto

# Docker build
docker-build:
	docker build -t my-service .

# Help target to display available targets
help:
	@echo "Available targets:"
	@echo "  make          : Build the binary (default target)"
	@echo "  make build    : Build the binary"
	@echo "  make clean    : Clean up the binary"
	@echo "  make deps     : Install dependencies"
	@echo "  make test     : Run tests"
	@echo "  make mod      : Update and tidy dependencies"
	@echo "  make run      : Build and run the application"
	@echo "  make docker-build  : Build Docker image"
	@echo "  make generate-proto : Generate Go code from proto files"
	@echo "  make help     : Show this help message"

# PHONY targets
.PHONY: default build clean deps test mod run docker-build help


