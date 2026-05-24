GOPATH := $(shell go env GOPATH)

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint:
	$(GOPATH)/bin/golangci-lint run ./...

lint-fix:
	$(GOPATH)/bin/golangci-lint run --fix ./...
