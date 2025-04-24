# GOEXPERT: Open Telemetry
Este projeto é uma API desenvolvida em Go que recebe um CEP como entrada, consulta a localização correspondente utilizando a API ViaCEP e retorna as informações climáticas atuais da cidade utilizando a WeatherAPI. As temperaturas são exibidas em três formatos: Celsius, Fahrenheit e Kelvin.

A aplicação está preparada para ser executada localmente ou ser implantada em ambiente serverless via Google Cloud Run, com suporte a tracing distribuído via OpenTelemetry, Zipkin, Jaeger, Prometheus e Grafana.

[![Go](https://img.shields.io/badge/go-1.23-informational?logo=go)](https://go.dev)

## Como rodar o projeto: make
```shell
## 1. Crie o .env
cp .env.example .env

## 2. Coloque sua api-key como valor na variável OPEN_WEATHERMAP_SERVICE_API_KEY no .env

## 3. Rode o comando
make run

## 4. Avalie os valores de retorno
make check
```

## Como rodar o projeto: manual
```shell
 ## 1. Gerar o build do container
docker-compose up -d --build

## 2. Avalie os valores de retorno
echo -n "422: "; curl -s "http://127.0.0.1:8080/get_address/1234567"
echo -n "404: "; curl -s "http://127.0.0.1:8080/get_address/12345678"
echo -n "200: "; curl -s "http://127.0.0.1:8080/get_address/13330250"

```
## Requisitos - Serviço A:
* O sistema deve receber um input de 8 dígitos via POST, através do schema:  { "cep": "29902555" }
* O sistema deve validar se o input é valido (contem 8 dígitos) e é uma STRING
* Caso seja válido, será encaminhado para o Serviço B via HTTP

- Caso não seja válido, deve retornar:
    - *Código HTTP:* **422**
    - *Mensagem:* **invalid zipcode**

---


## Requisitos - Serviço B:
* O sistema deve receber um CEP válido de 8 digitos
* O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin juntamente com o nome da localização.
* O sistema deve responder adequadamente nos seguintes cenários:

- Em caso de sucesso:
    - *Código HTTP:* 200
    - *Response Body:*
    ```json 
    {
        "city": "Salvador",
        "temp_C": 28.2,
        "temp_F": 82.76,
        "temp_K": 301.2
    } 
    ```
- Em caso de falha, caso o CEP não seja válido (com formato correto):
    - *Código HTTP:* 422
    - *Mensagem:* **invalid zipcode**
  

- Em caso de falha, caso o CEP não seja encontrado:
    - *Código HTTP:* 404
    - *Mensagem:* **can not find zipcode**

---

## OTEL + Zipkin
Após a implementação dos serviços, adicione a implementação do OTEL + Zipkin:
- Implementar tracing distribuído entre **Serviço A** - **Serviço B**
- Utilizar span para medir o tempo de resposta do serviço de busca de CEP e busca de temperatura
---

## Autor
**Leonardo Pinho**
- [GitHub](https://github.com/leonardopinho)
- Email: [contato@leonardopinho.com](mailto:contato@leonardopinho.com)
