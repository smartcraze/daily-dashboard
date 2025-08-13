FROM golang:1.22 AS builder

WORKDIR /app

# Download dependencies first (better cache)

COPY go.mod go.sum ./
RUN go mod download

# Copy source and build

COPY . .
RUN go build -o server .

# Runtime stage (small image)
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .
EXPOSE 8080

CMD ["./server"]
