# Dockerfile
FROM golang:1.23.1 AS builder

WORKDIR /app

# Copy go.mod and go.sum to the builder context
COPY ../go.mod ../go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire auth folder and pkg folder
COPY ../ ./

# Build the application
RUN go build -o ./ecom-micro-auth ./cmd

# Final stage
FROM alpine:3.20

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/ecom-micro-auth .

EXPOSE 8099

CMD ["./ecom-micro-auth"]
