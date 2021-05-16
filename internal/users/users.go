package users

import (
	"context"
	"database/sql"

	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
)

type UserServiceImpl struct {
	Db *sql.DB
}

func NewUserServiceImpl(db *sql.DB) UserService {
	return &UserServiceImpl{
		Db: db,
	}
}

func (u *UserServiceImpl) GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := models.Users().All(ctx, u.Db)
	return users, err
}
