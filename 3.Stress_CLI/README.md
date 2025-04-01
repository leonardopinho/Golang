# Stress test - FullCycle Go Expert Challenge

https://goexpert.fullcycle.com.br/pos-goexpert/

[![Go](https://img.shields.io/badge/go-1.22-informational?logo=go)](https://go.dev)

Projeto desenvolvido em **Go** com o objetivo de criar uma ferramenta de **teste de carga via linha de comando (CLI)**, permitindo avaliar o desempenho de serviços web. O sistema permite que o usuário defina a **URL alvo**, o **número total de requisições** e o nível de **concorrência** (quantidade de chamadas simultâneas), simulando cenários reais de alto tráfego. Após a execução, a ferramenta gera um **relatório completo**, informando o tempo total do teste, o número de requisições bem-sucedidas (status 200), e a distribuição dos demais códigos HTTP, como 404 ou 500. 

## Funcionalidades
1. Execução de testes de carga via linha de comando (CLI).
2. Definição de parâmetros de teste através de flags:
    - `--url`: URL do serviço a ser testado.
    - `--requests`: Quantidade total de requisições.
    - `--concurrency`: Número de requisições simultâneas.
3. Envio de requisições HTTP distribuídas conforme o nível de concorrência definido.
5. Geração automática de relatório ao final do teste com as seguintes informações:
    - Tempo total da execução.
    - Total de requisições realizadas.
    - Quantidade de respostas com status HTTP 200.
    - Distribuição de demais códigos HTTP (ex: 404, 500).

---

### Download dependencies

```
$ go mod tidy
```
