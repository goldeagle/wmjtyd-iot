# 交叉编译设置
TARGET_OS = linux
TARGET_ARCH = amd64
CC = /usr/bin/gcc
CXX = /usr/bin/g++
LD = /usr/bin/ld

APP_NAME := wmjtyd-iot
GO_FILES := $(shell find . -name '*.go')
GO_PACKAGES := $(shell go list ./...)

.PHONY: daemon build clean test fmt run install-deps lint help

daemon: ## Build the application's backend daemon only
	@echo "Building $(APP_NAME) for $(TARGET_OS)-$(TARGET_ARCH)..."
	env GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) CC=$(CC) CXX=$(CXX) LD=$(LD) go build -o $(APP_NAME) .
	chmod a+x $(APP_NAME)

build: ## Build the application
	@echo "Building $(APP_NAME) for $(TARGET_OS)-$(TARGET_ARCH)..."
	env GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) CC=$(CC) CXX=$(CXX) LD=$(LD) go build -o $(APP_NAME)_$(TARGET_OS)-$(TARGET_ARCH) .
	chmod a+x $(APP_NAME)_$(TARGET_OS)-$(TARGET_ARCH)

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f $(APP_NAME)
	rm -f $(APP_NAME)_$(TARGET_OS)-$(TARGET_ARCH)

test: ## Run tests
	@echo "Running tests..."
	go test -v $(GO_PACKAGES)

fmt: ## Format code
	@echo "Formatting code..."
	go fmt $(GO_PACKAGES)

run: build ## Run the application
	@echo "Running $(APP_NAME)..."
	./$(APP_NAME)

install-deps: ## Install dependencies
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

lint: ## Run static analysis
	@echo "Running static analysis..."
	golangci-lint run

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 --name $(APP_NAME) $(APP_NAME)

docker-push: ## Push Docker image to registry
	@echo "Pushing Docker image..."
	@echo "Please set your registry info before using this command"

docker-compose-up: ## Start dependent services with docker-compose
	@echo "Starting dependent services..."
	docker-compose up -d

docker-compose-down: ## Stop dependent services
	@echo "Stopping dependent services..."
	docker-compose down

docker-compose-logs: ## View service logs
	@echo "Showing service logs..."
	docker-compose logs -f

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
