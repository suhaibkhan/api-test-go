# Define the Golang builder stage
FROM golang:1.22.1-alpine3.19 AS builder

# Set working directory
WORKDIR /app

# Download dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Define the final stage
FROM alpine:3.19

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main main

# Expose the port (replace 8080 with your service port)
EXPOSE 8080

# Run the Go binary as the entrypoint
CMD ["./main"]
