# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23.1 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /pismo-app ./cmd/app/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:3.20.3
RUN apk update && apk add ca-certificates && apk add tzdata

WORKDIR /

COPY --from=build-stage /pismo-app /pismo-app

EXPOSE 8080

ENTRYPOINT ["/pismo-app"]
