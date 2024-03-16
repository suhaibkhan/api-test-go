BINARY_NAME=api_test_go.out

build:
	go build -o bin/${BINARY_NAME} main.go

run:
	go run main.go

