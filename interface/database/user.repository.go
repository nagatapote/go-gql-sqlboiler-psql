package database

import (
	"context"
	"go-gql-sqlboiler-psql/domain/models"
	"go-gql-sqlboiler-psql/infrastructure/db"
)

type userRepositoryImpl struct {
	DbUtils db.DbUtils
}

type UserRepository interface {
	FetchAll(context.Context) ([]*models.User, error)
}

func NewUserRepository(dbUtils db.DbUtils) UserRepository {
	return &userRepositoryImpl{
		DbUtils: dbUtils,
	}
}

func (r *userRepositoryImpl) FetchAll(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, r.DbUtils.GetDao(ctx))
}