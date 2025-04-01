# Rate Limiter - FullCycle Go Expert Challenge

https://goexpert.fullcycle.com.br/pos-goexpert/

[![Go](https://img.shields.io/badge/go-1.22-informational?logo=go)](https://go.dev)

Projeto desenvolvido em **Go** para implementar um *rate limiter* configurável, capaz de controlar o tráfego de requisições por **endereço IP** ou **token de acesso**, com prioridade para os limites definidos por token. Utilizando **Redis** como mecanismo de persistência, o sistema permite definir regras personalizadas, como o número máximo de requisições por segundo e o tempo de bloqueio, tudo isso configurável via variáveis de ambiente ou arquivo `.env`. Implementado como middleware, o rate limiter responde com status `429` ao exceder os limites e foi projetado com uma estratégia desacoplada para facilitar a troca do mecanismo de armazenamento no futuro.


---
### Funcionalidades

- Controle de requisições por IP ou token de acesso, com prioridade para os limites definidos pelo token.
- Definição de limites personalizáveis de requisições por segundo.
- Configuração de tempo de bloqueio ao exceder o limite de requisições.
- Utiliza Redis como mecanismo de persistência de dados.
- Estrutura de "strategy" para facilitar a substituição do Redis por outro sistema de armazenamento.


---

### Download dependencies

```
$ go mod tidy
```


### Execução

Para iniciar o servidor, utilize o comando:
```bash
docker build -t stress_cli .
```

### Testes

```
go test ./...
```

---

## Autor

**Leonardo Pinho**
- [GitHub](github.com/leonardopinho)
- Email: [contato@leonardopinho.com](mailto:contato@leonardopinho.com)
- Aluno da Pós **GoExpert 2024** pela [FullCycle](https://fullcycle.com.br)