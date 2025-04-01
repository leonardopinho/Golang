FROM golang:1.22-alpine AS builder
WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rate-limit main.go

FROM scratch

COPY --from=builder /app/.env .
COPY --from=builder /app/rate-limit .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["./rate-limit"]