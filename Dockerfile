# Build Stage
FROM golang:1.24.2 as builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

# Run Stage
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]
