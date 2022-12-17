package database

import (
	"context"
	"go-gql-sqlboiler-psql/domain/models"
	"go-gql-sqlboiler-psql/infrastructure/db"
	"go-gql-sqlboiler-psql/usecase/repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepositoryImpl struct {
	DbUtils db.DbUtils
}

func NewUserRepository(dbUtils db.DbUtils) repository.UserRepository {
	return &userRepositoryImpl{
		DbUtils: dbUtils,
	}
}

func (r *userRepositoryImpl) FetchAll(ctx context.Context) ([]*models.User, error) {
	results, err := models.Users().All(ctx, r.DbUtils.GetDao(ctx))
	return results, r.DbUtils.Error(err)
}

func (r *userRepositoryImpl) Fetch(ctx context.Context, id int64) (*models.User, error) {
	result, err := models.Users(
		qm.Where("id = ? and deleted_at is null", id),
	).One(ctx, r.DbUtils.GetDao(ctx))
	return result, r.DbUtils.Error(err)
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	return r.DbUtils.Error(user.Insert(ctx, r.DbUtils.GetDao(ctx), boil.Infer()))
}