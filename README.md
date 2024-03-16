## Run

1. Run postgres DB
2. Update postgres connection string in `internal/repository/init.go`
3. `go run main.go`

## Build binary and run

1. `go build -o api_test_go.out`
2. `./api_test_go.out`

## Build and run using Docker

1. `docker build -t api-test-go .`
2. `docker run -p 8080:8080 api-test-go`
