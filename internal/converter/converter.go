package converter

import (
	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
)

func mapToDBUser(u *model.NewUser) *models.User {
	return &models.User{
		Email:    u.Email,
		Username: u.Username,
	}
}

