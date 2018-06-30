all: build

help:
	@echo test generate build run query

test:
	go vet ./...
	golint ./...
	go test ./...

generate:
	go generate .

build: 
	go build .
	ls -l

run: build
	./gowebasm
