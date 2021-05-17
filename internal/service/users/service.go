package users

import (
	"context"

	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
)

type UserService interface {
	FindUserById(ctx context.Context, ID int64) (*models.User, error)
	GetUsers(context.Context) ([]*models.User, error)
	CreateUser(ctx context.Context, user model.NewUser) (*models.User, error)
	UpdateUser(ctx context.Context, user model.UpdateUser) (*models.User, error)
	DeleteUser(ctx context.Context, id int64) error
}
