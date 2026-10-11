.PHONY: build test lint lint-fix format format-check

GOLANGCI_LINT ?= golangci-lint
GOPATH := $(shell go env GOPATH)

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

test:
	go test ./...

lint:
	$(GOLANGCI_LINT) run ./...

lint-fix:
	$(GOLANGCI_LINT) run --fix ./...

format:
	$(GOLANGCI_LINT) fmt ./...

format-check:
	$(GOLANGCI_LINT) fmt --diff ./...
