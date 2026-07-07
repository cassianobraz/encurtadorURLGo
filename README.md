# Encurtador de URL em Go

API HTTP simples para encurtar URLs e consultar a URL original por um codigo curto.
O projeto usa Go, `net/http`, roteador `chi` e Redis para persistir os dados.

## Funcionalidades

- Cria codigos aleatorios de 8 caracteres para URLs enviadas via JSON.
- Consulta a URL original a partir de um codigo curto.
- Persiste os codigos no Redis usando um hash chamado `encurtador`.
- Inclui middlewares de recuperacao de panic, request ID e log de requisicoes.

## Requisitos

- Go `1.26.2` ou versao compativel com o `go.mod`.
- Docker e Docker Compose para subir o Redis localmente.

## Como executar

Na raiz do projeto, suba o Redis:

```bash
docker compose up -d
```

Depois execute a API:

```bash
go run ./cmd/api/main.go
```

O servidor sobe em:

```text
http://localhost:8080
```

Por padrao, a aplicacao conecta no Redis em:

```text
localhost:6379
```

## Endpoints

### Criar URL encurtada

```http
POST /api/url/shorten
Content-Type: application/json
```

Corpo da requisicao:

```json
{
  "url": "https://google.com"
}
```

Resposta de sucesso:

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "data": {
    "code": "eTgpLYzq"
  }
}
```

O valor em `data.code` e o codigo que deve ser usado para consultar a URL original.

### Consultar URL original

```http
GET /api/url/{code}
```

Exemplo:

```bash
curl -i http://localhost:8080/api/url/eTgpLYzq
```

Quando o codigo existe, a API retorna a URL original:

```json
{
  "data": {
    "full_url": "https://google.com"
  }
}
```

Quando o codigo nao existe, retorna `404 Not Found`.

## Exemplos com curl

Criar uma URL encurtada:

```bash
curl -X POST http://localhost:8080/api/url/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'
```

Consultar a URL original:

```bash
curl -i http://localhost:8080/api/url/SEU_CODIGO
```

Tambem ha um arquivo [`server.http`](server.http) com requisicoes prontas para clientes HTTP compativeis, como a extensao REST Client do VS Code.

## Estrutura do projeto

```text
.
|-- cmd/
|   `-- api/
|       `-- main.go              # Inicializacao do servidor HTTP e conexao com Redis
|-- internal/
|   |-- api/
|   |   |-- api.go               # Rotas, middlewares e resposta JSON
|   |   |-- get_shortened_url.go # Handler para consultar URLs
|   |   `-- shorten_url.go       # Handler para criar codigos curtos
|   `-- store/
|       |-- gen_code.go          # Geracao dos codigos curtos
|       `-- store.go             # Persistencia e consulta no Redis
|-- compose.yaml                 # Redis local via Docker Compose
|-- go.mod                       # Modulo e dependencias
|-- go.sum                       # Checksums das dependencias
|-- server.http                  # Exemplos de requisicao
`-- README.md
```

## Observacoes

- Os links encurtados sao persistidos no Redis e continuam disponiveis enquanto o volume Docker `redis` existir.
- O servidor escuta sempre na porta `8080`.
- O codigo curto e gerado com letras maiusculas, minusculas e numeros.
