FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server cmd/server/main.go


# Этап 2: Финальный минималистичный образ
FROM alpine:3.19

# Добавляем часовые пояса и сертификаты (нужны для S3/HTTPS)
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/config.toml .

EXPOSE 8080
CMD ["./server"]