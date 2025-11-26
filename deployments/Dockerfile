# Builder stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/api/main.go

# Run stage
FROM scratch
COPY --from=builder /app/bin/server /server
COPY --from=builder /app/config /config

# TODO: .env
CMD ["/server", "--config", "config/production.yaml"]
