# Build stage
FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o task-manager ./cmd/server

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/task-manager .
EXPOSE 8080
CMD ["./task-manager"]