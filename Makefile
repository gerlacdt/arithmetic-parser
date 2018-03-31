.PHONY= build start test test-no-cache


build:
	go build -o app cmd/main.go

start: build
	./app

test:
	go test ./...

test-no-cache:
	go test ./... -count 1
