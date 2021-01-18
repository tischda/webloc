# ---------------------------------------------------------------------------
# Makefile for GO utilities
# ---------------------------------------------------------------------------

PROJECT_DIR=$(notdir $(shell pwd))

BUILD_TAG=$(shell git describe --tags)
LDFLAGS=all=-ldflags "-X main.version=${BUILD_TAG} -s -w"

all: get build

build:
	go build -mod vendor ${LDFLAGS}

get:
	go get

test: fmt vet
	go test -mod vendor -v -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt

vet:
	go vet -v

install:
	go install ${LDFLAGS} ./...

dist: clean build
	upx -9 ${PROJECT_DIR}.exe

clean:
	go clean
