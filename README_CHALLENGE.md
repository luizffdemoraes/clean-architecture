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
| **Execução automática** | Ao rodar `docker compose up`: o banco sobe, as migrações são aplicadas automaticamente e a aplicação inicia nas portas configuradas. |

> **Atenção:** Garanta que a aplicação **aguarde o banco estar pronto** antes de rodar migrações ou iniciar (tratamento de *race condition* na inicialização).

## Como executar

Comando único:

```bash
docker compose up
```

### Portas dos serviços

| Serviço | Porta |
|---------|-------|
| **Web (REST)** | _[documente a porta, ex.: `8080`]_ |
| **gRPC** | _[documente a porta, ex.: `50051`]_ |
| **GraphQL** | _[documente a porta, ex.: `8081`]_ |

_Ajuste a tabela acima para refletir exatamente as portas definidas no seu `docker-compose.yaml`._

## Arquivos auxiliares

### `api.http`

Na **raiz** do repositório, crie um arquivo **`api.http`** com requisições prontas para:

1. **Criar** uma Order (para popular o banco e testar).
2. **Listar** as Orders (para validar o desafio).

## Entregável

| Item | Descrição |
|------|-----------|
| **Link do repositório** | [Inclua aqui o link do seu repositório no GitHub](https://github.com/seu-usuario/seu-repositorio) |
| **README** | Este arquivo, com o comando único de execução e as portas de Web, gRPC e GraphQL. |

## Regras de entrega

### Repositório exclusivo (muito importante)

Este repositório deve conter **apenas** o código deste desafio.

- **Não** entregue um repositório *monorepo* com pastas de outros cursos ou desafios anteriores — isso **bloqueia** o processo de correção automática.

### Branch principal

Todo o código deve estar na branch **`main`**.
