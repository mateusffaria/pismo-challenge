## Docker setup

- docker build -t <image-tag> -f ./deployments/Dockerfile .
- swagger setup
  go install github.com/go-swagger/go-swagger/cmd/swagger@latest
  swag init -g ./cmd/app/main.go -o ./docs

## Challenge backlog

Basic setup:

- [] Docker setup
- [] Swagger setup
- [] Adjust Readme containning all instructions (run, docker setup, testing)
- [] Git/Github setup
- [] .env setup

Feature checklist:

- [] Automated tests (unit, integration and E-E)
- []
