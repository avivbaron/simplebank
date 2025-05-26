# Build Stage
FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o sb main.go

# Run Stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/sb .
COPY app.env .

EXPOSE 8080
CMD ["/app/sb"]