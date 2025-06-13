FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o auth-service .

FROM alpine

WORKDIR /app
COPY --from=builder /app/auth-service .

COPY migrations ./migrations

EXPOSE 8080
CMD ["./auth-service"]