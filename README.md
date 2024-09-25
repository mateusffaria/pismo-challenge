# Pismo challenge

This project is a Go application designed with a focus on maintainability, scalability, and modern software architecture principles.

## Prerequisites

Ensure the following dependencies are installed:

- **Go**: version 1.23.1
- **Docker**: version 27.2.1
- **Docker Compose**
- **migrate**: database migration tool
- **Swagger**: API documentation tools

## Running the Application Locally (Without Docker)

To run the application locally, ensure all dependencies are installed and properly configured:

```bash
go mod tidy   # Install Go dependencies
go run ./cmd/app
```

## Docker Setup (Local Environment)

To spin up the entire application in a Dockerized environment, run the following command:

```bash
docker compose --env-file=./configs/.local.env -f ./deployments/docker-compose.yaml up app
```

This will bring up the app with its necessary services defined in the `docker-compose.yaml` file.

## API Documentation (Swagger)

The API documentation is auto-generated using Swagger. Ensure the Swagger tools are installed:

```bash
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate Swagger docs:

```bash
swag init -g ./cmd/app/main.go -o ./docs
```

Access the generated documentation via:

[Swagger UI](http://localhost:8080/api/docs/index.html)

## Running Tests

### Setting up the Test Database

To run tests, ensure the test database is up:

```bash
docker compose --env-file=./configs/.local.env -f ./deployments/docker-compose.yaml up db_test
```

### Unit Tests

Run unit tests for the application:

```bash
go test -v ./internal/...
```

### Integration Tests

Before running integration tests, make sure the test database is running (`db_test`):

```bash
go test -v ./test/...
```

## Database Migrations

Manage database schema changes using `golang-migrate`. To create a new migration:

```bash
migrate create -ext sql -dir db/migrations -seq create_operation_types_table
```

Ensure the `migrate` tool is installed and available in your `PATH`.

## Live Reload for Development

For hot-reloading during development, use [Air](https://github.com/cosmtrek/air):

```bash
go install github.com/cosmtrek/air@latest
air
```

This will automatically reload the application on file changes, enhancing development productivity.

## Key Features

- **Go Project Standard**: Follows best practices in Go application structure.
- **Hexagonal Architecture**: Implements domain-driven design principles.
- **Automated Testing**: Full suite of unit and integration tests.
- **Dockerized**: All services run in Docker containers for easy local setup.
- **Database Migrations**: Handles schema changes using `golang-migrate`.
- **Request Validation**: Uses `go-playground/validator` for robust validation.
- **Live Reload**: The `Air` package ensures seamless reloading during development.

## Swagger Documentation

Full API specification is available [here](https://github.com/mateusffaria/pismo-challenge/blob/main/docs/swagger.yaml).
