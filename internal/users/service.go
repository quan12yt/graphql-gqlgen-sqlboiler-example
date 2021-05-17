package users

import (
	"context"

	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
)

type UserService interface {
	FindUserById(ctx context.Context, ID int64) (*models.User, error)
	GetUsers(context.Context) ([]*models.User, error)
	CreateUser(ctx context.Context, user *model.NewUser) error
}
