package godevusecase

import (
	"context"
	"fmt"
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

func (u *userUseCaseImpl) Fetch(ctx context.Context, id int64) (*graphql.UserDetail, error) {
	m, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	} else if m == nil {
			return nil, fmt.Errorf("user not found id = %d", id)
	}
	return u.converter.UserModelToUserDetail(m)
}

func(u *userUseCaseImpl) Create(ctx context.Context, input graphql.UserCreateInput) (*graphql.UserDetail, error) {
	m, err := u.converter.UserCreateInputToUserModel(input)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Create(ctx, m); err != nil {
		return nil, err
	}
	return u.converter.UserModelToUserDetail(m)
}

func(u *userUseCaseImpl) Update(ctx context.Context, input graphql.UserCreateInput) (*graphql.UserDetail, error) {
	m, err := u.converter.UserCreateInputToUserModel(input)
	if err != nil {
		return nil, err
	}
	columns := u.converter.ConvertUpdateInputToDBColumnNames()
	if err := u.repository.Update(ctx, m, columns); err != nil {
		return nil, err
	}
	return u.converter.UserModelToUserDetail(m)
}