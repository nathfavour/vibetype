.PHONY: build run clean

BINARY_NAME=vibetype
BUILD_DIR=build

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)
