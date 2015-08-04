#
# Makefile for GO utilites
# 
# Compiler: GO 1.4.2
# 

build: fmt
	go build -ldflags "-X main.version `git describe --tags` -s"

fmt:
	go fmt

test:
	go test -v -cover

install:
	go install -ldflags "-X main.version `git describe --tags` -s"

clean:
	go clean
