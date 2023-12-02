all: fmt lint test build

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...

build:
	@mkdir -p bin
	go build -o ./bin ./... 
