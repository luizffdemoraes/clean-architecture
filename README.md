# Clean Architecture: Listagem de Orders (REST, gRPC e GraphQL)

## Objetivo

Neste desafio, você deve implementar a funcionalidade de **Listagem de Orders** em sua aplicação Clean Architecture. O objetivo principal é provar o **desacoplamento da arquitetura**: você criará um único Use Case (`ListOrders`) e o exporá através de **três interfaces de comunicação diferentes** simultaneamente.

## Tecnologias e padrões

| Aspecto | Tecnologia |
|---------|------------|
| **Linguagem** | Go (Golang) |
| **Arquitetura** | Clean Architecture |
| **Comunicação** | REST, gRPC e GraphQL |
| **Infraestrutura** | Docker e Docker Compose |

## Requisitos técnicos

### Use Case

- Crie o caso de uso de listagem de pedidos: **`ListOrdersUseCase`**.

### Interfaces de entrada

Disponibilize o acesso a esse Use Case através de:

| Interface | Especificação |
|-----------|----------------|
| **REST** | `GET /order` |
| **gRPC** | Service `ListOrders` |
| **GraphQL** | Query `ListOrders` |

### Banco de dados

- Crie as **migrações** necessárias para criar as tabelas do banco de dados.
- O banco deve ser **provisionado via Docker**.

## Requisitos de dockerização (automação total)

O avaliador **não deve executar nenhum comando manual** além do Docker Compose.

| Item | Descrição |
|------|-----------|
| **Container da aplicação** | Dockerfile para a aplicação Go. |
| **Orquestração** | `docker-compose.yaml` sobe o banco de dados e o container da aplicação. |
| **Execução automática** | Ao rodar `docker compose up --build`: o banco sobe, as migrações são aplicadas automaticamente e a aplicação inicia nas portas configuradas. |

> **Atenção:** Garanta que a aplicação **aguarde o banco estar pronto** antes de rodar migrações ou iniciar (tratamento de *race condition* na inicialização).

## Como executar

Comando único (recomendado para garantir a imagem da aplicação atualizada):

```bash
docker compose up --build
```

Para executar em background:

```bash
docker compose up -d --build
```

Para parar os containers:

```bash
docker compose down
```

Para acompanhar logs da aplicação:

```bash
docker compose logs -f app
```

## Criação de tabelas (migração)

Nao existe processo manual para criar tabela neste projeto.

- A tabela `orders` e criada automaticamente no startup da aplicacao;
- o script `scripts/entrypoint.sh` espera o MySQL ficar pronto e executa todos os arquivos `.sql` de `migrations/`;
- a migracao atual esta em `migrations/001_create_orders.sql`.

Se quiser recriar tudo do zero (incluindo volume do banco), use:

```bash
docker compose down -v && docker compose up --build
```

### Portas dos serviços

| Serviço | Porta |
|---------|-------|
| **Web (REST)** | `8000` |
| **gRPC** | `50051` |
| **GraphQL** | `8080` |

## Arquivos auxiliares

### `api.http`

Na **raiz** do repositório, crie um arquivo **`api.http`** com requisições prontas para:

1. **Criar** uma Order (para popular o banco e testar).
2. **Listar** as Orders (para validar o desafio).

### Pasta `docs`

Para facilitar os testes, existe uma pasta `docs` com:

- `docs/clean-architecture-orders.postman_collection.json` (collection Postman: REST, GraphQL e pasta **gRPC (grpcurl)** com comandos para `ListOrders` e `CreateOrder`)

## Como testar rapidamente

> Antes dos testes, garanta que a stack esteja em execucao com `docker compose up --build` (ou `docker compose up -d --build`).

### Teste REST (Create + List)

Criar order:

```bash
curl --request POST \
  --url http://localhost:8000/order \
  --header 'Content-Type: application/json' \
  --data '{
  "id": "order-1",
  "price": 100.5,
  "tax": 0.5
}'
```

Listar orders:

```bash
curl --request GET \
  --url http://localhost:8000/order
```

### Teste GraphQL (Query ListOrders)

```bash
curl --request POST \
  --url http://localhost:8080/query \
  --header 'Content-Type: application/json' \
  --data '{"query":"query { ListOrders { id Price Tax FinalPrice } }"}'
```

Playground GraphQL:

- `http://localhost:8080/`

### Teste gRPC (ListOrders e CreateOrder)

Com **`grpcurl` instalado** (por exemplo `go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest`):

Listar orders:

```bash
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```

Criar order:

```bash
grpcurl -plaintext \
  -d '{"id":"order-2","price":200.0,"tax":10.0}' \
  localhost:50051 pb.OrderService/CreateOrder
```

**Sem instalar o `grpcurl` no sistema**, use a imagem oficial (Docker baixa na primeira vez). Em Linux, `--network host` faz o container enxergar `localhost:50051`:

Listar orders:

```bash
docker run --rm --network host fullstorydev/grpcurl:latest \
  -plaintext localhost:50051 pb.OrderService/ListOrders
```

Criar order:

```bash
docker run --rm --network host fullstorydev/grpcurl:latest \
  -plaintext \
  -d '{"id":"order-2","price":200.0,"tax":10.0}' \
  localhost:50051 pb.OrderService/CreateOrder
```

**Opcional — inspecionar a API (server reflection)**  
Úteis para ver assinaturas e métodos sem abrir o `.proto`. Com `grpcurl` local, use o mesmo padrão trocando `docker run ...` por `grpcurl` e removendo `--network host`.

Listar métodos do serviço:

```bash
docker run --rm --network host fullstorydev/grpcurl:latest -plaintext \
  localhost:50051 list pb.OrderService
```

Descrever o serviço ou um método:

```bash
docker run --rm --network host fullstorydev/grpcurl:latest -plaintext \
  localhost:50051 describe pb.OrderService

docker run --rm --network host fullstorydev/grpcurl:latest -plaintext \
  localhost:50051 describe pb.OrderService.CreateOrder

docker run --rm --network host fullstorydev/grpcurl:latest -plaintext \
  localhost:50051 describe pb.OrderService.ListOrders
```

> **Nota:** em cada novo `CreateOrder`, use um `id` de string **único**; repetir o mesmo `id` gera erro de chave duplicada no MySQL.

## Entregável

| Item | Descrição |
|------|-----------|
| **Link do repositório** | [https://github.com/luizffdemoraes/clean-architecture](https://github.com/luizffdemoraes/clean-architecture) |
| **README** | Este arquivo, com o comando único de execução e as portas de Web, gRPC e GraphQL. |

## Regras de entrega

### Repositório exclusivo (muito importante)

Este repositório deve conter **apenas** o código deste desafio.

- **Não** entregue um repositório *monorepo* com pastas de outros cursos ou desafios anteriores — isso **bloqueia** o processo de correção automática.

### Branch principal

Todo o código deve estar na branch **`main`**.
