package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

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

func (r *queryResolver) FindByID(ctx context.Context, id string) (*models.User, error) {
	id1, _ := strconv.ParseInt(id, 10, 64)
	return r.sv.FindUserById(ctx, id1)
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
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Users returns generated.UsersResolver implementation.
func (r *Resolver) Users() generated.UsersResolver { return &usersResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type usersResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func NewSchemaConfig(us users.UserService) generated.Config {
	return generated.Config{
		Resolvers: NewResolver(us),
	}
}
func NewResolver(us users.UserService) *Resolver {
	return &Resolver{us}
}
