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

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
