# Build Stage
FROM golang:1.24.3 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .


FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .
COPY wait-for-postgres.sh .
RUN apt-get update && apt-get install -y netcat-openbsd && chmod +x wait-for-postgres.sh



CMD ["./wait-for-postgres.sh", "./main"]
