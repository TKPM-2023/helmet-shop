ifneq (,$(wildcard ./.env))
    include .env
    export
endif


dev:
	go run ./cmd/main.go

build:
	rm ./build-out || true
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build-out cmd/main.go
	upx -9 -q ./build-out