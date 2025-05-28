# Build Stage
FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app

# Install curl for downloading migrate
RUN apk add --no-cache curl

# Copy source code
COPY . .

# Build Go binary
RUN go build -o main main.go

# Download and extract golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
  | tar -xz

# Run Stage
FROM alpine:3.21
WORKDIR /app

# Copy built binary and migrate CLI
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate

# Copy environment and migrations
COPY app.env .
COPY db/migration ./migration

# Copy and start script
COPY start.sh ./start.sh
RUN chmod +x ./start.sh

# Copy and start wait-for script
COPY wait-for.sh ./wait-for.sh
RUN chmod +x ./start.sh ./wait-for.sh

EXPOSE 8080

ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]