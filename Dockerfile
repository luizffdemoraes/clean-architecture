FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/ordersystem ./cmd/ordersystem

FROM alpine:3.22

WORKDIR /app
RUN apk add --no-cache mysql-client

COPY --from=builder /bin/ordersystem /app/ordersystem
COPY migrations /app/migrations
COPY scripts/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

EXPOSE 8000 50051 8080
ENTRYPOINT ["/app/entrypoint.sh"]
