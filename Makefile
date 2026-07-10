GOPATH := $(shell go env GOPATH)

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

test:
	go test ./...

lint:
	$(GOPATH)/bin/golangci-lint run ./...

lint-fix:
	$(GOPATH)/bin/golangci-lint run --fix ./...

