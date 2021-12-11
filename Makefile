.DEFAULT_GOAL := build
export integration=true
tools:
		go install github.com/goblinfactory/gf-markdown
.PHONY:tools

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet: fmt lint
		go vet ./...
.PHONY:vet

build: vet tools
		go test ./...
		gf-markdown readme.md
		go build 		
.PHONY:build
