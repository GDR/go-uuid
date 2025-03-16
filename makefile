BINARY_NAME=uuidserver
SRC_DIR=./
BUILD_DIR=./build
GO_FILES=$(wildcard $(SRC_DIR)*.go)

all: build

build: $(GO_FILES)
	@echo "Building the Go application..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

run: build
	@echo "Running the application..."
	GIN_MODE=release $(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...
	