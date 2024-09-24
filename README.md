## Docker setup

- docker build -t <image-tag> -f ./deployments/Dockerfile .
- docker compose --env-file=./configs/.local.env -f ./deployments/docker-compose.yaml up
- swagger setup
  go install github.com/go-swagger/go-swagger/cmd/swagger@latest
  swag init -g ./cmd/app/main.go -o ./docs

## Challenge backlog

Basic setup:

- [] Docker setup
- [x] Swagger setup
- [] Adjust Readme containning all instructions (run, docker setup, testing)
- [x] Git/Github setup
- [x] .env setup
- [] golang-lint

Feature checklist:

- [x] Persistance layer
- [] Automated tests (unit, integration and E-E)
- [x] Create User Account
  - [x] Unit test
  - [x] Implementation
  - [] Refactor -> Remove "User" from method names
  - [x] Error handling
- [x] Get User Account
  - [x] Unit test
  - [x] Implementation
  - [x] Refactor
  - [x] Error handling
- [] Create User Transaction
  - [] Unit test
  - [] Integration test
  - [] Implementation
    - [] Enum for transaction type validation
  - [] Refactor
  - [] Error handling
