#
# Makefile for GO utilites
# 
# Compiler: GO 1.5
# 

build: get
	go build -ldflags "-X main.version=`git describe --tags` -s"

get:
	go get

test: fmt
	go test -v -cover

fmt:
	go fmt

install:
	go install -a -ldflags "-X main.version=`git describe --tags` -s"

clean:
	go clean
