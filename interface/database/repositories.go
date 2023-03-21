package database

import (
	"go-gql-sqlboiler-psql/infrastructure/db"
	"go-gql-sqlboiler-psql/usecase/repository"
)

type AppRepositories struct {
	UserRepository repository.UserRepository
}

func NewAppRepositories(dbUtil db.DbUtils) *AppRepositories {
	return &AppRepositories{
		UserRepository: NewUserRepository(dbUtil),
	}
}
