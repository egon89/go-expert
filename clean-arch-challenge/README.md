# Clean Arch Challenge

## Application
To run:
```bash
# create a .env file
cp .env.example .env

# run the application
make run

# run the application in development mode
make dev-run
```

## Database
Access docker database:
```bash
docker-compose exec mysql bash
mysql -u root -p

show databases;
use orders;
show tables;
```

### golang-migrate
[Installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#linux-deb-package)

```bash
# create a migration
migrate create -ext=sql -dir=database/migrations -seq init

# up
migrate -path=database/migrations -database="mysql://user:password@tcp(localhost:3306)/orders" -verbose up

# down
migrate -path=database/migrations -database="mysql://user:password@tcp(localhost:3306)/orders" -verbose up
```

Or you can just run the following sql into database:
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

### Playground
Go to [http://localhost:8080](http://localhost:8080) in your browser to access the GraphQL Playground

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

### Configuration
- Change the paths in *gqlgen.yml* to generate file in desired folder (internal/infra/graph)
- Use the file *schema.graphqls* to create the schemas
  - use the **Mutation** type to do some action in the types
  - use the **Query** type to get informations
- **Resolvers** are "methods" that will be executed when a GraphQL action is called

The command `go run github.com/99designs/gqlgen generate`
will change the *schema.resolvers.go* file according to *schema.graphqls* file

### Errors
Invalid memory address or nil pointer error when try generate command:
```bash
rm go.sum && go get -u github.com/99designs/gqlgen
go mod tidy

go run github.com/99designs/gqlgen generate
```

## Google Wire
[Documentation](https://github.com/google/wire)

```bash
go install github.com/google/wire/cmd/wire@latest

# check if the wire folder was created
ls ~/go/bin
```

After change the *wire.go* file, run the `wire` command into *cmd/ordersystem* folder to update the *wire_gen.go* file.


## gRPC
[Go documentation](https://grpc.io/docs/languages/go/quickstart)

Install and check the protobuf compiler version
```bash
sudo apt install protobuf-compiler

protoc --version
```

Install Go plugins
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# to check if the folders protoc-gen-go and protoc-gen-go were created
ls ~/go/bin
```

Install __vscode-proto3__ plugin to work with `.proto` files.

After change a *.proto* file run the following command in the root project folder:
```bash
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```

The files *order_grpc.pb.go* and *order.pb.go* in *pb* folder will be updated.


### Install Evans
```bash
go install github.com/ktr0731/evans@latest

# to check if the folder evans was created
ls ~/go/bin
```

Run **evans** command in the project root folder:
```bash
evans --proto internal/infra/grpc/protofiles/order.proto repl
```

<p>Show available services:</p>
<img src="docs/evans-show-service.png">

<p>Create an order:</p>
<img src="docs/evans-create-order.png">

<p>Get all orders:</p>
<img src="docs/evans-get-all-orders.png">


## Golang configuration
Check if PATH variable has golang configuration or add in the `~/.zshrc` file the following line:
```
export PATH="$HOME/go/bin:/usr/local/go/bin:$PATH"
```
