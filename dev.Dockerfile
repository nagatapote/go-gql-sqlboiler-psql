FROM golang:1.17.0-alpine
ENV GO111MODULE=on

WORKDIR /go/go-gql-sqlboiler-psql
RUN apk update && apk add git alpine-sdk

COPY . .

RUN go mod download

RUN go install github.com/cosmtrek/air@v1.29.0 \
  && go install github.com/99designs/gqlgen@v0.14.0 \
  && go install github.com/rubenv/sql-migrate/...@v1.1.1 \
  && go install github.com/volatiletech/sqlboiler/v4@latest && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

EXPOSE 8000