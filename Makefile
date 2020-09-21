.PHONY: build
build:
	go build -o bin/x-bot cmd/x-bot/main.go

.PHONY: format
format:
	go mod tidy
	goimports -w .

.PHONY: test
test:
	go test -v ./...
