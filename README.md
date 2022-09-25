# go-gql-sqlboiler-psql

go 、graphql、sqlboiler、psql 検証用

# 実行

```
docker compose build

make create-docker-network

make migration/new FILE_NAME=${filename}

make migrate/up

make gql

make up
```
