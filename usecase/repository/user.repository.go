package repository

import (
	"context"
	"go-gql-sqlboiler-psql/domain/models"
)

type UserRepository interface {
	FetchAll(context.Context) ([]*models.User, error)
	Fetch(context.Context, int64) (*models.User, error)
	Create(context.Context, *models.User) error
	Update(context.Context, *models.User, []string) error
	Delete(context.Context, int64) error
}