.PHONY= build start test test-no-cache


build:
	go build -o app cmd/main.go

start: build
	./app

test:
	go test ./... -v

test-no-cache:
	go test ./... -v -count 1
