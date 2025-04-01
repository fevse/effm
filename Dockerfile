FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .

RUN go mod download
RUN go build -o effm cmd/effm/main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/effm .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations/20250401102138_effm_stor.sql .

CMD ["./effm"]