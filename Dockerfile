FROM golang:1.17.0-alpine as builder
ENV GO111MODULE on

WORKDIR /go/go-gql-sqlboiler-psql

COPY go.mod .
COPY go.sum .

RUN apk --update --no-cache add git alpine-sdk
RUN set -x \
  && go mod download \
  && go install github.com/99designs/gqlgen@v0.14.0

COPY . .
RUN gqlgen && go build -o /server ./server.go

FROM alpine:3.15.3

EXPOSE 8000
RUN apk --update --no-cache add ca-certificates tzdata
COPY --from=builder ./server ./app/server
USER nobody
CMD [ "./app/server" ]