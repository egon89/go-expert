# Clean Arch Challenge

## Database
```bash
docker-compose exec mysql bash
mysql -u root -p
```

```sql
CREATE DATABASE IF NOT EXISTS orders;

USE orders;

CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) NOT NULL,
    price DOUBLE NOT NULL,
    tax DOUBLE NOT NULL,
    final_price DOUBLE NOT NULL,
    PRIMARY KEY (id)
);
```

## GraphQL
- change the paths in **gqlgen.yml** to generate file in desired folder (internal/infra/graph)

## Playground
```graphql
mutation {
  createOrder(input: { id: "123", Price: 100.0, Tax: 10.0 }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

```graphql
query {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### Errors
Invalid memory address or nil pointer error when try generate command
Fix:
```bash
rm go.sum && go get -u github.com/99designs/gqlgen
go mod tidy

go run github.com/99designs/gqlgen generate
```
