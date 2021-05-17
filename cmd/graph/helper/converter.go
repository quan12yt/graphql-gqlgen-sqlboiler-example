package helper

import (
	"fmt"

	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
)

func mapUser(u *models.User) *model.User {
	i := fmt.Sprintf("%d", u.ID)
	return &model.User{
		ID:       i,
		Email:    u.Email,
		Username: u.Username,
	}
}

func mapListUser(u []*models.User) []*model.User {
	users := make([]*model.User, len(u))
	for v, i := range u {
		users[v] = mapUser(i)
	}
	return users
}
