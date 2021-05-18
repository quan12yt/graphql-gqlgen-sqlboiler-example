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
	"github.com/quan12yt/graphql-sqlboiler-example/internal/service/users"
	log "github.com/sirupsen/logrus"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	createdUser, err := r.sv.CreateUser(ctx, input)

	if err != nil {
		return nil, err
	}
	return mapUser(createdUser), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	createdUser, err := r.sv.UpdateUser(ctx, input)

	if err != nil {
		return nil, err
	}
	return mapUser(createdUser), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	i, _ := strconv.ParseInt(id, 10, 64)
	err := r.sv.DeleteUser(ctx, i)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.sv.GetUsers(ctx)
	return mapListUsers(users), err
}

func (r *queryResolver) FindByID(ctx context.Context, id string) (*model.User, error) {
	i, _ := strconv.ParseInt(id, 10, 64)
	got, meetups, err := r.sv.FindUserById(ctx, i)
	log.Println(meetups[0])

	if err != nil {
		return nil, err
	}
	user := mapUser(got)
	user.Meetups = mapListMeetups(meetups)
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
func mapUser(u *models.User) *model.User {
	i := fmt.Sprintf("%d", u.ID)
	return &model.User{
		ID:       i,
		Email:    u.Email,
		Username: u.Username,
	}
}
func mapListUsers(u []*models.User) []*model.User {
	users := make([]*model.User, len(u))
	for v, i := range u {
		users[v] = mapUser(i)
	}
	return users
}
func mapListMeetups(m []*models.Meetup) []*model.Meetup {
	meetups := make([]*model.Meetup, len(m))

	for _, i := range m {
		id := fmt.Sprintf("%d", i.ID)
		me := &model.Meetup{
			ID:          id,
			Description: i.Description,
			Names:       i.Names,
		}
		meetups = append(meetups, me)
	}
	return meetups
}
