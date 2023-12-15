.DEFAULT_GOAL := install

install:
	go mod tidy

run-dev:
	air

run:
	go run server.go

build:
	#mkdir dist/
	go build -o dist/ .

