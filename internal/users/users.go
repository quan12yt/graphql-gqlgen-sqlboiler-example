package users

import (
	"context"
	"database/sql"

	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
	"github.com/volatiletech/sqlboiler/boil"
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

func (u *UserServiceImpl) FindUserById(ctx context.Context, ID int64) (*models.User, error) {
	user, err := models.FindUser(ctx, u.Db, ID)
	return user, err
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, user *model.NewUser) error {
	us := converter.mapToDBUser(user)
	return us.Insert(ctx, u.Db, boil.Infer())
}
