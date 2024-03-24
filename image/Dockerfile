FROM golang:1.22-alpine3.19


RUN umask 000

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/spf13/cobra-cli@latest


CMD air

