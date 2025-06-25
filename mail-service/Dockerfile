FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
COPY .env .
RUN go mod tidy
RUN go build -o mail-service ./cmd/main.go


FROM alpine

WORKDIR /app
COPY --from=builder /app/mail-service .
COPY .env .

EXPOSE 8081
CMD ["./mail-service"]