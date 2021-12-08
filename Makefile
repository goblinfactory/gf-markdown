.DEFAULT_GOAL := build
export integration=true
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
		go test ./...
		# go test
		markdown readme.md
		go build 
		
.PHONY:build
