package godevusecase

import (
	"context"
	"go-gql-sqlboiler-psql/domain/models/graphql"
	"go-gql-sqlboiler-psql/usecase/converter"
	"go-gql-sqlboiler-psql/usecase/repository"
)

type userUseCaseImpl struct {
	repository         repository.UserRepository
	converter          converter.UserConverter
}

type UserUseCase interface {
	FetchAll(context.Context) ([]*graphql.UserDetail, error)
	Fetch(context.Context, int64) (*graphql.UserDetail, error)
	Create(context.Context, *graphql.UserDetail) error
	Update(context.Context, *graphql.UserDetail, []string) error
	Delete(context.Context, int64) error
}

func NewUserUseCase(repository repository.UserRepository, converter converter.UserConverter) UserUseCase {
	return &userUseCaseImpl{
		repository:         repository,
		converter:          converter,
	}
}

func (u *userUseCaseImpl) FetchAll(ctx context.Context) ([]*graphql.UserDetail, error) {
	ms, err := u.repository.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return u.converter.UserModelsToUserDetails(ms)
}