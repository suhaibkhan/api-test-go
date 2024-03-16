## Run

Run postgres DB
Update postgres connection string in `internal/repository/init.go`
go run main.go

## Build binary and run

go build -o api_test_go.out
./api_test_go.out

## Build and run using Docker

docker build -t api-test-go .
docker run -p 8080:8080 api-test-go
