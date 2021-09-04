.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: build
build: deps
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/lambda cmd/lambda/main.go

test:
	export AWS_REGION=ap-northeast-1;\
	export DYNAMO_ENDPOINT=http://localhost:4566;\
	export DYNAMO_TABLE_DEVICE_GPS=local_device_gps;\
	go test -race -v ./

testclean:
	go clean -testcache

run-local:
	export AWS_REGION=ap-northeast-1;\
	export DYNAMO_ENDPOINT=http://localhost:4566;\
	export DYNAMO_TABLE_DEVICE_GPS=local_device_gps;\
	go run cmd/gpsviewer/main.go --host 0.0.0.0 --port 3000


zip: build
	zip -j bin/lambda.zip bin/lambda