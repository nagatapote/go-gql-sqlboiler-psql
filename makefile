RUN=docker compose run --rm api

create-docker-network:
	docker network create go-dev-api-network

up:
	docker compose up db api

gql:
	${RUN} gqlgen

db-up:
	docker compose up db -d

db-down:
	docker compose down db

migration/new:
	${RUN} sql-migrate new --env="local" ${FILE_NAME}

migrate/up:
	make db-up
	sleep 5
	${RUN} sql-migrate up --env="local"

migrate/down:
	make db-up
	sleep 5
	${RUN} sql-migrate down --env="local"

sqlboiler:
	${RUN} sqlboiler psql

# schemaのlinter（引数がcamlCaseになっているかのチェック）
schema-linter:
	graphql-schema-linter --rules input-object-values-are-camel-cased  ./schema/*.graphql