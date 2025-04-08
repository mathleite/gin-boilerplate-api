FROM golang:alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o main ./

# Build a minimal image
FROM alpine:latest

# Add a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app

COPY --from=builder /app/main /usr/local/bin/main

EXPOSE 8000

CMD ["main"]
