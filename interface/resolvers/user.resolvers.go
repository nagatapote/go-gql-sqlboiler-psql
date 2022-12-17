package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	graphql1 "go-gql-sqlboiler-psql/domain/models/graphql"
	graphql2 "go-gql-sqlboiler-psql/infrastructure/graphql"
)

func (r *mutationResolver) UserCreate(ctx context.Context, params graphql1.UserCreateInput) (*graphql1.UserDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UserUpdate(ctx context.Context, params graphql1.UserUpdateInput) (*graphql1.UserDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UserDelete(ctx context.Context, id int64) (*graphql1.ClientDeleteResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*graphql1.UserDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int64) (*graphql1.UserDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns graphql2.MutationResolver implementation.
func (r *Resolver) Mutation() graphql2.MutationResolver { return &mutationResolver{r} }

// Query returns graphql2.QueryResolver implementation.
func (r *Resolver) Query() graphql2.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
