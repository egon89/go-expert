# client-server-api-challenge

## Creating a database
```shell
sqlite3 /home/user/sqlite-db/go-exchange.db
```

## Creating a table
Run the following command into database:

```sql
CREATE TABLE IF NOT EXISTS exchange (
    code TEXT,
    bid REAL,
    high REAL,
    low REAL,
    var_bid REAL,
    pct_change REAL,
    ask REAL,
    timestamp INTEGER,
    create_date DATETIME
);
```

## Creating and configure the .env file
```shell
cp .env.example .env
```


## Server
```shell
go run server.go
```
