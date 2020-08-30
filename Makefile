## Build all binaries 
build:
	$ go build -o bin/bookstore-golang-api internals/main.go
.PHONY: build

## Run development server
run:
	$ go run main.go
.PHONY: run'

## Start compiled app
start:
	$ sh -c './bin/bookstore-golang-api'
.PHONY: build
