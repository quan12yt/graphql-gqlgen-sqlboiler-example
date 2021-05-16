package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/generated"
	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.sv.GetUsers(ctx)
}

func (r *usersResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *usersResolver) Password(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *usersResolver) Active(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		r:  r,
		sv: r.sv,
	}
}

// Users returns generated.UsersResolver implementation.
func (r *Resolver) Users() generated.UsersResolver { return &usersResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct {
	r  *Resolver
	sv users.UserService
}
type usersResolver struct{ *Resolver }

func NewSchemaConfig(us users.UserService) generated.Config {
	return generated.Config{
		Resolvers: NewResolver(us),
	}
}
func NewResolver(us users.UserService) *Resolver {
	return &Resolver{us}
}
