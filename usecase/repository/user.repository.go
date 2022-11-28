package repository

import (
	"context"
	"go-gql-sqlboiler-psql/domain/models"
)

type UserRepository interface {
	FetchAll(context.Context) ([]*models.User, error)
}