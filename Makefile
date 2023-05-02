BIN_FILE=qsgo
BIN_DIR=bin
BIN_PATH=$(BIN_DIR)/$(BIN_FILE)

tidy:
	go mod tidy

run:
	go build -o $(BIN_PATH).exe && $(BIN_PATH).exe

default: windows

windows:
	@echo "Building for Windows..."
	go fmt
	go build -o $(BIN_PATH).exe
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"

