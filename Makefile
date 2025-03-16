# Project Name
BINARY_NAME := kernal_gpt

.PHONY: all
all: build

.PHONY: build
build: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: ensure-output-dir
ensure-output-dir:
	mkdir -p ./dist

.PHONY: build-linux-amd64
build-linux-amd64: ensure-output-dir
	GOOS=linux GOARCH=amd64 go build -o ./dist/$(BINARY_NAME)-linux-amd64 ./main.go

.PHONY: build-linux-arm64
build-linux-arm64: ensure-output-dir
	GOOS=linux GOARCH=arm64 go build -o ./dist/$(BINARY_NAME)-linux-arm64 ./main.go

.PHONY: build-darwin-amd64
build-darwin-amd64: ensure-output-dir
	GOOS=darwin GOARCH=amd64 go build -o ./dist/$(BINARY_NAME)-darwin-amd64 ./main.go

.PHONY: build-darwin-arm64
build-darwin-arm64: ensure-output-dir
	GOOS=darwin GOARCH=arm64 go build -o ./dist/$(BINARY_NAME)-darwin-arm64 ./main.go