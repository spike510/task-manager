# Build stage
FROM golang:1.24.5 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server ./cmd/server

# Run stage
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]