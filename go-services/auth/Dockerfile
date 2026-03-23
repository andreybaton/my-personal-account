FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go-services/auth/go.mod go-services/auth/go.sum ./
RUN go mod download
COPY go-services/auth/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/auth

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/auth-service .
EXPOSE 8080
CMD ["./auth-service"]