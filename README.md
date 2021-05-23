# Go Rest API
An API written in Go, with Mux for routing and PostgreSQL as database.

The example endpoint is `/products`. The structure of a product is:

```json
{
  "id": "6a9b6882-2838-4da2-ac90-83419f146135",
  "name": "Keychron K2",
  "price": 79.99,
  "description": "Keychron K2 is a 75% layout (84-keys) white led backlit compact Bluetooth mechanical keyboard.",
  "created_at": "2021-05-23T23:14:29.517606Z"
}
```

## Dependencies

- Go v1.16+
- Docker
- Docker-compose v1.29+
- [Migrate](https://github.com/golang-migrate/migrate)

## How To Run


### Postgres
- Create Postgres container with: `make postgres`
- To drop the container: `make postgresdown`

### Tests
- For running tests: `make test`

### Development Environment
- Create development database with: `make devdbup`
- Run server with: `go run main.go`


## Functionalities:

- [x] Database package tests
- [x] Database connection and repositories
- [x] Models package tests
- [x] Product models
- [x] HTTP Handlers functions
- [ ] HTTP Handlers tests
