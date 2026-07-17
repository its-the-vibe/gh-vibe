.PHONY: build test vet

build:
	go build

test:
	go test -v ./...

vet:
	go vet ./...
