FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/app
RUN go build -o notifications_service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cmd/app/notifications_service .

EXPOSE 8001

CMD ./notifications_service
