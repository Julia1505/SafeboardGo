.DEFAULT_GOAL := run

fmt: deps
	go fmt ./...
.PHONY:fmt lint vet clean build test coverage run

lint: fmt
	glint ./...

vet: fmt
	go vet ./...

clean:
	rm -rf ./bin/safeboard ./pkg/*/*.out ./pkg/*.out ./*.out

build: vet clean
	go build -o bin/safeboard cmd/safeboard/main.go

run: build
	./bin/safeboard

test: build
	go test -v ./pkg/decoder ./pkg/file ./pkg/parser

coverage:build
	go test -coverprofile=cover.out ./pkg/decoder ./pkg/file ./pkg/parser
	go tool cover -html=cover.out -o cover.html

deps:
	go get ./cmd/safeboard ./pkg/*

