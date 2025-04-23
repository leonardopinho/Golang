FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates

COPY zip_code_service/. .
COPY .env /app/.env

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o zip_code_service

FROM scratch
WORKDIR /app/project_a

#WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env /app

COPY --from=builder /app/zip_code_service .

ENTRYPOINT ["./zip_code_service"]
