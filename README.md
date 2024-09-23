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
- [] .env setup

Feature checklist:

- [] Persistance layer
- [] Automated tests (unit, integration and E-E)
- []
