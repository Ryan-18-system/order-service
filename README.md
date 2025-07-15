# Desafio - Listagem de Pedidos

## Como executar

```bash
docker compose up
```

Isso irá subir:

- PostgreSQL na porta `5432`
- REST API na porta `8080`
- gRPC na porta `50051`
- GraphQL na porta `8081`

## Endpoints

- REST: `GET /order`
- gRPC: `ListOrders` (proto em `/proto/order.proto`)
- GraphQL: `query { listOrders { id customer_name total created_at } }`

## Testes

Utilize o arquivo `api.http` para testar os endpoints REST.