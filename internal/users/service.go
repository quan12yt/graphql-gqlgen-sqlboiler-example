package users

import (
	"context"

	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
)

type UserService interface {
	GetUsers(context.Context) ([]*models.User, error)
}
