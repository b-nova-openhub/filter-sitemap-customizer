all: tidy build run

init:
	go mod download github.com/gorilla/mux
	go mod download github.com/spf13/cobra
	go fmt
	go mod tidy
	go mod vendor
	mkdir bin
	build
	run

tidy:
	go mod tidy
	go fmt ./...

build:
	go test ./...
	go build -o bin/fisicus main.go

run:
	chmod +x bin/fisicus
	chmod +x fisicus.sh
	./fisicus.sh

install:
	go install -v ./...