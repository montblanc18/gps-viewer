.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: build
build: deps
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/lambda cmd/lambda/main.go