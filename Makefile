all: build run

initall: init build run

init:
	go mod download github.com/gorilla/mux
	go mod download github.com/spf13/cobra
	go fmt
	go mod tidy
	go mod vendor
	mkdir bin

tidy:

build:
	go build -o bin/fisicus main.go

run:
	chmod +x bin/fisicus
	chmod +x fisicus.sh
	./fisicus.sh

install:
	go install -v ./...