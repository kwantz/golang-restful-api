# Golang RESTful API
Simple RESTful API created with Golang, MySQL, HTTP Router, Validation, and OpenAPI.

## Features (CRUD)
- Create Category
- Update Category
- Delete Category
- Get Category
- Get List Categories

## To Do
- Create OpenAPI (file: [openapi.yaml](https://github.com/kwantz/golang-restful-api/blob/master/openapi.yaml))
- Setup Database (file: [app/database.go](https://github.com/kwantz/golang-restful-api/blob/master/app/database.go), goto: [Setup](https://github.com/kwantz/golang-restful-api#setup))
- Create Entity (folder: [model/entity](https://github.com/kwantz/golang-restful-api/blob/master/model/entity))
- Create Repository (folder: [repository](https://github.com/kwantz/golang-restful-api/blob/master/repository))
- Create Service (folder: [service](https://github.com/kwantz/golang-restful-api/blob/master/service))
- Create Validation
- Create Controller (folder: [controller](https://github.com/kwantz/golang-restful-api/blob/master/service))
- HTTP Router (file: [app/router.go](https://github.com/kwantz/golang-restful-api/blob/master/app/router.go))
- HTTP Server (file: [main.go](https://github.com/kwantz/golang-restful-api/blob/master/main.go))
- Error Handler
- Authentication (file: [middleware/auth_middleware.go](https://github.com/kwantz/golang-restful-api/blob/master/middleware/auth_middleware.go))
- Integration Test / Benchmark (folder: [test](https://github.com/kwantz/golang-restful-api/blob/master/test))

## Setup
```bash
# Run docker
docker-compose up -d

# Create table database
cat sql/databases.sql | docker exec -i golang-restful-api-mysql /usr/bin/mysql -u root --password=password
cat sql/categories.sql | docker exec -i golang-restful-api-mysql /usr/bin/mysql -u root --password=password golang_restful_api
cat sql/categories.sql | docker exec -i golang-restful-api-mysql /usr/bin/mysql -u root --password=password golang_restful_api_test

# Run server
go run main.go

# Run testing
go test ./test/

# Run benchmark
go test -bench=. ./test/
```
