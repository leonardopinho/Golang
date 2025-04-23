.PHONY: tests build run server check

tests:
	@echo "Rodando testes..."
	go test ./... -v

build:
	@docker build -t weather_api .

run:
	@docker run --rm -p 8080:8080 --name weather_api_container weather_api

server:
	@go run main.go

check:
	@echo -n "422: "; curl -s "http://127.0.0.1:8080/get_weather?cep=1234567"
	@echo -n "404: "; curl -s "http://127.0.0.1:8080/get_weather?cep=12345678"
	@echo -n "200: "; curl -s "http://127.0.0.1:8080/get_weather?cep=13330250"