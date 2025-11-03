.PHONY: run build test clean dev

# Binary name
BINARY_NAME=tabloom

# Run application
run:
	@echo "Running Tabloom..."
	@go run cmd/tabloom/main.go

# Build application
build:
	@echo "Building Tabloom..."
	@go build -o bin/$(BINARY_NAME) cmd/tabloom/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download

# Help
help:
	@echo "Available commands:"
	@echo "  run    - Run application"
	@echo "  build  - Build binary"
	@echo "  test   - Run tests"
	@echo "  clean  - Clean binaries"
	@echo "  deps   - Install dependencies"