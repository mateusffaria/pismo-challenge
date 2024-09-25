## Docker setup

- docker build -t <image-tag> -f ./deployments/Dockerfile .
- docker compose --env-file=./configs/.local.env -f ./deployments/docker-compose.yaml up
- swagger setup
  go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go install github.com/swaggo/swag/cmd/swag@latest
  swag init -g ./cmd/app/main.go -o ./docs

test db: docker compose --env-file=./configs/.local.env -f ./deployments/docker-compose.yaml up postgres_test
migration creation: migrate create -ext sql -dir db/migrations -seq create_operation_types_table
separate unit/integartion:

test: go test -v ./...

live reload: air

## Challenge backlog

Basic setup:

- [x] Docker setup
- [x] Swagger setup
- [x] Adjust Readme containning all instructions (run, docker setup, testing)
- [x] Git/Github setup
- [x] .env setup
- [] golang-lint

Feature checklist:

- [x] Persistance layer
- [x] Automated tests (unit, integration and E-E)
- [x] Create User Account
  - [x] Unit test
  - [x] Implementation
  - [x] Refactor -> Remove "User" from method names
  - [x] Error handling
- [x] Get User Account
  - [x] Unit test
  - [x] Implementation
  - [x] Refactor
  - [x] Error handling
- [x] Create User Transaction
  - [x] Unit test
  - [x] Integration test
  - [x] Implementation
    - [x] Enum for transaction type validation
  - [x] Refactor
  - [x] Error handling
