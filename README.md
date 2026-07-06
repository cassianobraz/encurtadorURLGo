# Encurtador de URL em Go

API HTTP simples para encurtar URLs e redirecionar acessos por um codigo curto.
O projeto usa Go, `net/http` e o roteador `chi`.

## Funcionalidades

- Cria codigos aleatorios de 8 caracteres para URLs enviadas via JSON.
- Redireciona `GET /{code}` para a URL original.
- Mantem os dados em memoria usando um `map[string]string`.
- Inclui middlewares de recuperacao de panic, request ID e log de requisicoes.

## Requisitos

- Go `1.26.2` ou versao compativel com o `go.mod`.

## Como executar

Na raiz do projeto:

```bash
go run .
```

O servidor sobe em:

```text
http://localhost:8080
```

## Endpoints

### Criar URL encurtada

```http
POST /api/shorten
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
  "data": "eTgpLYzq"
}
```

O valor em `data` e o codigo que deve ser usado para acessar a URL encurtada.

### Redirecionar pela URL encurtada

```http
GET /{code}
```

Exemplo:

```bash
curl -i http://localhost:8080/eTgpLYzq
```

Quando o codigo existe, a API responde com redirecionamento permanente para a URL original.
Quando o codigo nao existe, retorna `404 Not Found`.

## Exemplos com curl

Criar uma URL encurtada:

```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'
```

Acessar a URL encurtada:

```bash
curl -i http://localhost:8080/SEU_CODIGO
```

Tambem ha um arquivo [`server.http`](server.http) com requisicoes prontas para clientes HTTP compativeis, como a extensao REST Client do VS Code.

## Estrutura do projeto

```text
.
|-- api/
|   `-- api.go       # Rotas, handlers, resposta JSON e geracao dos codigos
|-- main.go          # Inicializacao do servidor HTTP
|-- go.mod           # Modulo e dependencias
|-- go.sum           # Checksums das dependencias
|-- server.http      # Exemplos de requisicao
`-- README.md
```

## Observacoes

- Os links encurtados nao sao persistidos em banco de dados. Ao reiniciar o servidor, todos os codigos gerados sao perdidos.
- O servidor escuta sempre na porta `8080`.
- O codigo curto e gerado com letras maiusculas, minusculas e numeros.
