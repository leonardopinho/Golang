FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates

COPY weather_service/. .
COPY .env /app/.env

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o weather_service

FROM scratch
WORKDIR /app/project_b

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env /app

COPY --from=builder /app/weather_service .

ENTRYPOINT ["./weather_service"]
