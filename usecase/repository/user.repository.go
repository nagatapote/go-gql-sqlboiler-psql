package repository

import (
	"context"
	"go-gql-sqlboiler-psql/domain/models"
)

type UserRepository interface {
	FetchAll(context.Context) ([]*models.User, error)
	Fetch(context.Context, int64) (*models.User, error)
}