VERSION := $(shell git tag | tail -n1)
BINARY_NAME = general-ledger-api
BINARY_VERSIONED = ${BINARY_NAME}-${VERSION}

start: clean build run

clean:
	rm -f ${BINARY_NAME}-*

build:
	go build -o ${BINARY_VERSIONED}

run:
	./${BINARY_VERSIONED}

test:
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic
	go tool cover -html=coverage.txt
	rm coverage.txt

.PHONY: start clean build run test
