BINARY_NAME=routesApp

build:
	 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ./cmd/routesApp

test:
	go test -v ./...

run:
	go run ./cmd/routesApp