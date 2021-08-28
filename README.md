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
- Setup Database (goto: [Setup](https://github.com/kwantz/golang-restful-api#setup), file: [app/database.go](https://github.com/kwantz/golang-restful-api/blob/master/app/database.go))
- Create Entity (folder: [model/entity](https://github.com/kwantz/golang-restful-api/blob/master/model/entity))
- Create Repository (folder: [repository](https://github.com/kwantz/golang-restful-api/blob/master/repository))
- Create Service (folder: [service](https://github.com/kwantz/golang-restful-api/blob/master/service))
- Create Validation
- Create Controller (folder: [controller](https://github.com/kwantz/golang-restful-api/blob/master/service))
- HTTP Router (file: [app/router.go](https://github.com/kwantz/golang-restful-api/blob/master/app/router.go))
- HTTP Server (file: [main.go](https://github.com/kwantz/golang-restful-api/blob/master/main.go))
- Error Handler
- Authentication
- Integration Test

## Setup
```bash
# Run docker
docker-compose up -d

# Create table database
cat db.sql | docker exec -i golang-restful-api-mysql /usr/bin/mysql -u root --password=password golang_restful_api

# Run server
go run main.go
```
