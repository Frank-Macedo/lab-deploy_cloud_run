# AppForecast

API simples para consultar temperatura por CEP (Brasil). Este repositório fornece uma API em Go que consulta o ViaCEP para obter endereço a partir do CEP e a API do WeatherAPI para obter a temperatura atual.

**Arquivos principais:**
- [go.mod](go.mod#L1) — declara a versão do Go e dependências.
- [cmd/app/main.go](cmd/app/main.go#L1) — ponto de entrada e rotas HTTP.
- [internal/api/handlers/temperature_handler.go](internal/api/handlers/temperature_handler.go#L1) — handler `/temperature/{cep}`.
- [internal/infrastructure/via_cep/client.go](internal/infrastructure/via_cep/client.go#L1) — cliente para ViaCEP.
- [internal/infrastructure/weather/client.go](internal/infrastructure/weather/client.go#L1) — cliente para WeatherAPI.

## Tecnologias

- Go 1.24.x
- Gorilla Mux (router)
- Docker (imagem para empacotamento)
- WeatherAPI (serviço externo)
- ViaCEP (serviço externo para CEP)

## Pré-requisitos

- Go 1.24
- Docker (opcional, para rodar em container)
- Chave da WeatherAPI (variável `API_KEY`)

## Instalação (local)

1. Clone o repositório:

```
git clone <repo-url>
cd AppForecast
```

2. Rodar em modo desenvolvimento:

```
go run ./cmd/app/main.go
```

Exemplo (rodando dentro de `cmd/app`):

```
export API_KEY=123
go run main.go
```


## Build e execução (binário)

```
go build -o app ./cmd/app/main.go
./app
```

## Docker

Build da imagem:

```
docker build -t appforecast:latest .
```

Run (exemplo definindo `API_KEY` e `PORT`):

```
docker run -e API_KEY="$API_KEY" -e PORT=8080 -p 8080:8080 appforecast:latest
```

Observação: o `Dockerfile` expõe a porta `8080` e a aplicação usa `8080` por padrão quando `PORT` não está definido.

## Docker Compose

O `compose.yaml` incluso expõe `8080` e pode injetar `API_KEY`. Se quiser expor outra porta, adicione `PORT` às variáveis de ambiente ou ajuste o mapeamento de portas.

Exemplo sugerido (trecho):

```
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - API_KEY=${API_KEY}
      - PORT=8080
```

## Variáveis de ambiente

- `API_KEY` — chave da WeatherAPI (obrigatória).

## Endpoints

-- `GET /` — Mensagem de boas-vindas.
  - Exemplo: `curl http://localhost:8080/` → `Welcome to the Weather API!`

- `GET /temperature/{cep}` — Recupera temperatura atual para o CEP (8 dígitos sem hífen).
  - Parâmetro: `cep` (ex.: `01310100`)
  - Resposta (200): texto JSON com temperatura em °C, °F e K, exemplo: `25°C, 77°F, 298.15K` (conteúdo retornado como `application/json`).
  - Erros:
    - `422 Unprocessable Entity` — CEP inválido (formato errado).
    - `404 Not Found` — CEP não encontrado ou erro na busca de clima.

Observação: O handler constrói a requisição consultando o ViaCEP para obter `localidade` e `uf`, e depois passa essa informação para a WeatherAPI.

## Testes

Rodar todos os testes:

```
go test ./...
```

Observações:
- Há testes unitários que usam mocks (por exemplo, `internal/application/usecases/get_temperature_usecase_test.go`).
- O teste de handler (`internal/api/handlers/temperature_handler_test.go`) faz chamadas que podem alcançar os serviços externos; para testes determinísticos recomendamos desabilitar rede ou mockar clients.

## Observações e TODOs

-- O `compose.yaml` atualmente seta `API_KEY`. Se você quiser forçar uma porta específica diferente da padrão `8080`, defina `PORT` em `compose.yaml` ou ajuste o mapeamento de portas.


