package converter

import (
	"go-gql-sqlboiler-psql/domain/models"
	"go-gql-sqlboiler-psql/domain/models/graphql"
	"go-gql-sqlboiler-psql/utils"
)

type userConverterImpl struct {
}

type UserConverter interface {
	UserModelsToUserDetails([]*models.User) ([]*graphql.UserDetail, error)
	UserModelToUserDetail(*models.User) (*graphql.UserDetail, error)
	UserCreateInputToUserModel(graphql.UserCreateInput) (*models.User, error)
	UserUpdateInputToUserModel(graphql.UserUpdateInput) (*models.User, error)
	ConvertUpdateInputToDBColumnNames() []string
}

func NewUserConverter() UserConverter {
	return &userConverterImpl{}
}

func (c *userConverterImpl) UserModelsToUserDetails(ms []*models.User) ([]*graphql.UserDetail, error) {
	var results []*graphql.UserDetail
	for _, m := range ms {
		result, err := c.UserModelToUserDetail(m)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (c *userConverterImpl) UserModelToUserDetail(m *models.User) (*graphql.UserDetail, error) {
	return &graphql.UserDetail{
		ID:                         m.ID,
		Name:                       m.Name,
		Email:                      m.Email,
	}, nil
}

func (c *userConverterImpl) UserCreateInputToUserModel(input graphql.UserCreateInput) (*models.User, error) {
	return &models.User{
		Name:                       input.Name,
		Email:                      input.Email,
	}, nil
}

func (c *userConverterImpl) UserUpdateInputToUserModel(input graphql.UserUpdateInput) (*models.User, error) {
	return &models.User{
		ID:                         input.ID,
		Name:                       input.Name,
		Email:                      input.Email,
	}, nil
}

func (c *userConverterImpl) ConvertUpdateInputToDBColumnNames() []string {
	m := models.User{}
	g := graphql.UserUpdateInput{}
	return utils.ConvertUpdateInputToDBColumnNames(m, g)
}