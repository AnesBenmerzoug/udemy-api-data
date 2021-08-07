BIN_DIR=bin/
BIN="$(BIN_DIR)/udemy-api-data"

.PHONY: all run fmt vet clean build


all: fmt vet

run:
	go run main.go

fmt:
	gofmt -s -w .

vet:
	go vet ./...

build:
	go build -o $(BIN) main.go

clean:
	go clean