package resolvers

import "go-gql-sqlboiler-psql/usecase/godevusecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUseCase godevusecase.UserUseCase
}
