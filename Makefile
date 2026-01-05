all: build

build:
	mkdir -p bin
	go build -o bin/goserverps .

run: build
	./bin/goserverps

run-dev:
	go run main.go

clean:
	rm -rf bin
	rm -rf .cache

fmt:
	gofmt -w .

vet:
	go vet ./...

.PHONY: all build run run-dev clean fmt vet
