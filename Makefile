.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet: fmt lint
		go vet ./...
.PHONY:vet

build: vet
		go test -tags integration ./...
		go build ./markdown.go
.PHONY:build
