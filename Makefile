.PHONY= build start test test-no-cache


build:
	go build -o evaluator cmd/evaluator.go

start: build
	./evaluator

test:
	go test ./...

test-no-cache:
	go test ./... -count 1
