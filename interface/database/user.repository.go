package database

import (
	"context"
	"database/sql"
	"go-gql-sqlboiler-psql/domain/models"
)

type userRepositoryImpl struct {
	Db *sql.DB
}

type UserRepository interface {
	FetchAll(context.Context) ([]*models.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		Db: db,
	}
}

func (r *userRepositoryImpl) FetchAll(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, r.Db)
}