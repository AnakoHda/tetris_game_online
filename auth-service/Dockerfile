FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o auth-service ./cmd/main.go


RUN wget -O migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.tar.gz && \
    mv migrate /usr/bin/migrate && \
    chmod +x /usr/bin/migrate


FROM alpine

WORKDIR /app
COPY --from=builder /app/auth-service .

COPY --from=builder /usr/bin/migrate /usr/bin/migrate
COPY ./migrations ./migrations
COPY .env .


EXPOSE 8080
CMD migrate -path ./migrations -database "$DB_URL" up && ./auth-service