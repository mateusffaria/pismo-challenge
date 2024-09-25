# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23.1 AS build-stage

WORKDIR /pismo-challenge

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /pismo-app ./cmd/app/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Create docs directory and ensure permissions
RUN mkdir -p /app/docs && chmod -R 777 /app/docs

# Run swag init with the correct paths and permissions
RUN /go/bin/swag init -g ./cmd/app/main.go -o ./docs --parseDependency --parseInternal
RUN go test -v ./internal/...

# EXPOSE 8080

# ENTRYPOINT ["/pismo-app"]

# Deploy the application binary into a lean image
FROM alpine:3.20.3
RUN apk update && apk add ca-certificates && apk add tzdata

WORKDIR /

COPY --from=build-stage /pismo-app /pismo-app
COPY --from=build-stage /pismo-challenge/db/migrations /db/migrations
COPY --from=build-stage /pismo-challenge/docs /docs

EXPOSE 8080

ENTRYPOINT ["/pismo-app"]
