TARGET = notes-server

SRC_DIR = ../cmd
BUILD_DIR = ./
GO = go

.PHONY: run build clean help

run: build
	./$(BUILD_DIR)/$(TARGET)

build: clean
	$(GO) mod tidy
	$(GO) build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(TARGET) $(SRC_DIR)

clean:
	rm -f $(BUILD_DIR)/$(TARGET)

help:
	@echo "Available commands:"
	@echo "	build:"
	@echo "		Build $(TARGET) in $(PWD)/$(BUILD_DIR)"
	@echo "	clean:"
	@echo "		Delete $(TARGET)"
	@echo "	help:"
	@echo "		Show this message"

.DEFAULT_GOAL := build