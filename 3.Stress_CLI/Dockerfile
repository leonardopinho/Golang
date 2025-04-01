FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk update && apk upgrade

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build -ldflags="-s -w" -o stress_cli .

FROM scratch

COPY --from=builder /app/stress_cli /usr/bin/stress_cli

ENTRYPOINT ["stress_cli"]