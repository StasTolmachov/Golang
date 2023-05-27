.PHONY: build

build:
	go build ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

.DEFAULT_GOAL := build
