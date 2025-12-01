# --- Stage 1: Builder ---
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git tzdata ca-certificates

ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/server ./cmd/api/main.go

# --- Stage 2: Final ---
FROM gcr.io/distroless/static-debian12

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /app/bin/server /server
COPY --from=builder /app/config /config

USER appuser:appuser

ENTRYPOINT ["/server"]
CMD ["--config", "config/production.yaml"]
