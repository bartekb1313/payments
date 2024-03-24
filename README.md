# Payment Monitoring

Simple web application for payment monitoring. It has dynamic panel and REST API.
Since I am using HTMX (https://htmx.org) no additional frontend application needs to be created.


## Stack/libs
- Go
- Docker
- PostgreSql
- HTMX (https://github.com/bigskysoftware/htmx)
- Oapi-codegen (https://github.com/deepmap/oapi-codegen)
- Cobra (https://github.com/spf13/cobra)
- Air for development (https://github.com/cosmtrek/air)
- Taskfile (https://taskfile.dev/)

## UI
- html templates from tabler.io (https://tabler.io)

## OpenApi Specification

- spec file: image/app/docs/openapi/payments.yml
- generated server layer: image/app/internal/common/server/spec/payments.gen.go

## Local development
You will need two things: Docker and Taskfile (https://taskfile.dev/)

Taskfile commands:
- task docker:start - start development environment
- task docker:logs - container logs
- task docker:sh - shell of development container
- task migrations:create - new empty migration file
- task migrations:up - apply migration
- task openapi:generate - regenerate API server layer
