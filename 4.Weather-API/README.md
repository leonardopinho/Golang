# GOEXPERT: weather-api
Projeto do treinamento GoExpert(FullCycle).

[![Go](https://img.shields.io/badge/go-1.22-informational?logo=go)](https://go.dev)


## Desafio
Desenvolvimento de uma API em Go capaz de receber um CEP, consultar a cidade correspondente (ViaCEP) e, em seguida, buscar os dados climáticos da localidade por meio da WeatherAPI. A resposta retorna a temperatura atual formatada em Celsius, Fahrenheit e Kelvin.

## Como rodar o projeto: make
```shell
# build the container image
make build

# run locally
make run
```


## Como rodar o projeto: manual
```shell
## 1. Crie o .env
cp .env.example .env

## 2. Coloque sua api-key como valor na variável OPEN_WEATHERMAP_SERVICE_API_KEY no .env

## 3. Gerar o build do container
docker build -t weather_api .

## 4. Rodar o projeto
docker run --rm -p 8080:8080 --name weather_api_container weather_api

## 5. Avalie os valores de retorno
echo -n "422: "; curl -s "http://127.0.0.1:8080/get_weather?cep=1234567"
echo -n "404: "; curl -s "http://127.0.0.1:8080/get_weather?cep=12345678"
echo -n "200: "; curl -s "http://127.0.0.1:8080/get_weather?cep=13330250"
```

## Requisitos:
- [x] O sistema deve receber um CEP válido de 8 digitos
- [x] O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- [x] O sistema deve responder adequadamente nos seguintes cenários:
    - Em caso de sucesso:
        - [x] Código HTTP: 200
        - [x] Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
    - Em caso de falha, caso o CEP não seja válido (com formato correto):
        - [x] Código HTTP: 422
        - [x] Mensagem: invalid zipcode
    - ​​​Em caso de falha, caso o CEP não seja encontrado:
        - [x] Código HTTP: 404
        - [x] Mensagem: can not find zipcode
- [x] Deverá ser realizado o [deploy no Google Cloud Run](https://cloudrun-goexpert-747247348579.us-central1.run.app/get_weather?cep=13330250).

## Autor
**Leonardo Pinho**
- [GitHub](https://github.com/leonardopinho)
- Email: [contato@leonardopinho.com](mailto:contato@leonardopinho.com)
