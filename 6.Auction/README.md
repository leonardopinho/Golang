# GOEXPERT: Auction
Projeto do laboratório "Concorrência com Golang - Leilão" do treinamento GoExpert (FullCycle).

[![Go](https://img.shields.io/badge/go-1.20-informational?logo=go)](https://go.dev)

## Descrição
A rotina original de criação de leilões e gerenciamento de lances já está implementada. A melhoria proposta envolve adicionar uma rotina automática que encerra os leilões que atingiram seu tempo limite.

Para atingir esse objetivo, utilizamos `goroutines` em Go para garantir a concorrência e eficiência da aplicação.

## Funcionalidades Implementadas
* **Cálculo automático do tempo do leilão**

  * O tempo de duração do leilão é definido por variáveis de ambiente.

* **Encerramento automático dos leilões vencidos**

  * Uma `goroutine` monitora continuamente leilões em aberto e fecha automaticamente os que ultrapassarem o tempo estipulado.

* **Testes Automatizados**

  * Teste unitário que garante o correto funcionamento do fechamento automático dos leilões.

## Como rodar o projeto: make
```shell
## 1. Inicie o projeto:
make run

## 2. Crie um registro que será fechado automaticamente:
make create

## 3. Liste os leilões criados:
make show

## 4. Aguarde o período configurado na variável AUCTION_DURATION definida no arquivo de configuração.
  
## 3. Verifique o fechamento automático:
make check  

## 4. Rodar teste:
make test
```

## Como rodar o projeto: manual
```shell
 ## 1. Gerar o build do container
docker-compose up --build

## 2. Avalie os valores de retorno usando o arquivo api.http
```


## Autor
**Leonardo Pinho**
- [GitHub](https://github.com/leonardopinho)
- Email: [contato@leonardopinho.com](mailto:contato@leonardopinho.com)
