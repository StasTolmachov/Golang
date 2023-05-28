.PHONY: build
build:
	go build ./cmd/apiserver

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test: tidy
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
