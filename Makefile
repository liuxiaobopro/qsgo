BIN_FILE=qsgo
BIN_DIR=.
BIN_PATH=$(BIN_DIR)/$(BIN_FILE)

tidy:
	@go mod tidy

build: windows

windows:
	@echo "Building for Windows..."
	@go fmt
	@go build -o $(BIN_PATH).exe
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"

linux:
	@echo "Building for Linux..."
	@go fmt
	@go build -o $(BIN_PATH)
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"

clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_PATH).exe

run: tidy clean windows
	@echo "Running..."
	@$(BIN_PATH).exe $(ARGS)