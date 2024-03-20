FROM golang:1.22-alpine AS builder
LABEL authors="qrave1"


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/app .
COPY --from=builder /app/.env .

CMD ["./app"]
