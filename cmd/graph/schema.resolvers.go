package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/generated"
	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.sv.GetUsers(ctx)
	return helper.mapListUsers(users), err
}

func (r *queryResolver) FindByID(ctx context.Context, id string) (*model.User, error) {
	i, _ := strconv.ParseInt(id, 10, 64)
	//(ctx, err)
	got, err := r.sv.FindUserById(ctx, i)
	//dieIF(ctx, err)
	user := helper.mapUser(got)
	return user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

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
