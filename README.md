# bbob
Short for bob the builder, BBob is a Database Intializer Written in Golang.

## Supported Databases
- [ ] MySQL
- [x] PostgreSQL
- [ ] SQLite
- [ ] MongoDB
- [x] Cassandra

## Installation

```bash
go get github.com/karim-w/bbob@latest
```

## Usage
```bash
bbob -config config.yaml
```

## Configuration
```yaml
# config.yaml
sql:
  - databaseDsn: "w/e"
    databaseType: postgres
    databaseNames:
      - bob_the_builder_postgres
cql:
  - urls:
      - "w/e"
    username: "w/e"
    password: "w/e"
    class: "SimpleStrategy"
    replicationFactor: 1
    keyspaces:
      - bob_the_builder_cassandra
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Authors 
- [karim-w]("https://github.com/karim-w")



