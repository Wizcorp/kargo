.PHONY: all fmt build

all: fmt build

build:
	@mkdir -p bin
	go build -o bin/kargo -v kargo/*.go

fmt:
	go fmt compose/*.go
	go fmt kargo/*.go
	go fmt machine/*.go
	go fmt utils/*.go
