
OUTPUT = onwallex
MAIN = cmd/main.go

.DEFAULT_GOAL := help

.PHONY: help lint build test clean 


help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'


lint:
	@echo run golangci lint 
	@golangci-lint run


build:
	@echo building $(OUTPUT) 
	@go build -o $(OUTPUT) $(MAIN)


test: 
	@echo run go test 
	@go test ./...


clean: 
	@rm -f $(OUTPUT)


