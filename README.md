# Sahabat Rental-StockMovement

This microservice works solely to Stock Movement.

## Requirements

- Go v1.20 and above
- Makefile (optional)
- MySQL

## Installation

`make install` if your IDE not automate it for you. Alternatively, just run `go mod tidy`.

> fill your `.env`.

## Development

`make run-dev` requires [Air](https://github.com/cosmtrek/air).

## Running

`go run server.go`

## Building

`make build` or `go build -o dist/ .`

> cross-platform compilation just need to put `GOOS` and `GOARCH`.