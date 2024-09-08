# Rate Limiter em Go

Este projeto implementa um rate limiter em Go que pode ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso. O rate limiter utiliza Redis para armazenar e gerenciar os dados de requisições e pode ser integrado como um middleware em um servidor web.

## Estrutura do Projeto

- `main.go` – Configuração do servidor web e inicialização do rate limiter.
- `middleware.go` – Middleware que aplica o rate limiter.
- `limiter.go` – Lógica do rate limiter.
- `redis.go` – Configuração e conexão com o Redis.
- `.env` – Arquivo de configuração com variáveis de ambiente.
- `docker-compose.yml` – Configuração do Docker para iniciar o Redis.

## Pré-requisitos

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (para desenvolvimento e execução)

## Configuração

### 1. Clonar o Repositório

Clone o repositório para sua máquina local:

```sh
git clone https://github.com/4lexRossi/go-rate-limiter.git
cd go-rate-limiter
```

### 2. Configurar o Redis com Docker

Para iniciar o Redis, execute o seguinte comando:

```sh
docker-compose up -d
```

Isso criará e iniciará um contêiner Redis. O Redis estará acessível em `localhost:6379`.

### 3. Configurar Variáveis de Ambiente

Existe um arquivo .env na pasta raiz do projeto com as seguintes variáveis:

```env
REDIS_ADDR=localhost:6379
RATE_LIMIT_IP=5
RATE_LIMIT_TOKEN=10
BLOCK_DURATION=300
```

`REDIS_ADDR` – Endereço do servidor Redis.

`RATE_LIMIT_IP` – Número máximo de requisições permitidas por segundo para um único IP.

`RATE_LIMIT_TOKEN` – Número máximo de requisições permitidas por segundo para um único token.

`BLOCK_DURATION` – Tempo em segundos durante o qual um IP ou token será bloqueado após exceder o limite.

### 4. Executar o Servidor
Para executar o servidor, use o seguinte comando:

```sh
go run main.go
```

O servidor web será iniciado na porta `8080`.

## Testes Automatizados
Os testes automatizados podem ser executados com o comando:

```sh
go test ./...
```

Certifique-se de que o Redis esteja em execução antes de executar os testes.

## Uso
O rate limiter é integrado como um middleware para o servidor web. Requisições com o header `API_KEY` serão limitadas com base no token especificado, e todas as requisições serão limitadas com base no IP do cliente.

### Exemplo de Requisição
Para testar o rate limiter, você pode usar um cliente HTTP como `curl` ou ferramentas como Postman. Adicione o header `API_KEY` com um valor de token e envie requisições para `http://localhost:8080`.

```sh
curl -H "API_KEY: <TOKEN>" http://localhost:8080
```
Se o limite for excedido, a resposta será um erro 429 com a mensagem:

```css
you have reached the maximum number of requests or actions allowed within a certain time frame
```

## Configuração do Rate Limiter
As configurações do rate limiter podem ser ajustadas modificando o arquivo .env. Reinicie o servidor após alterar as configurações para que as mudanças tenham efeito.

Contribuição
Se você encontrar problemas ou quiser melhorar o projeto, fique à vontade para abrir uma issue ou enviar um pull request.
