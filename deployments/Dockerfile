# Builder stage
FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git tzdata ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/api/main.go

# Run stage
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/bin/server /server
COPY --from=builder /app/config /config

CMD ["/server", "--config", "config/production.yaml"]
