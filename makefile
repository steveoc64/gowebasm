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
	cd server && go build .
	cd wasm && GOARCH=wasm GOOS=js go build -o code.wasm .
	ls -l server wasm

run: build
	./server/server
