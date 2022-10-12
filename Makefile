.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	glint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o bin/safeboard cmd/safeboard/main.go
.PHONY:build

run: build
	./bin/safeboard

test: build
	go test -v
.PHONY:test

coverage:build
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
.PHONY:coverage