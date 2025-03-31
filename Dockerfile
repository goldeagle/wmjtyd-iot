# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o wmjtyd-iot .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/wmjtyd-iot .
COPY config/config.yaml ./config/

# Expose HTTP port
EXPOSE 8080

# Run the application
CMD ["./wmjtyd-iot"]