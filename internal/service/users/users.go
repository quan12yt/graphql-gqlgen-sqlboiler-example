package users

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/model"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	if err != nil {
		return nil, errors.New("cant find any user for the given id")
	}
	return user, nil
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, user model.NewUser) (*models.User, error) {
	var a models.User
	a.Email = user.Email
	a.Username = user.Username
	err := a.Insert(ctx, u.Db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &a, nil

}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, user model.UpdateUser) (*models.User, error) {
	i, _ := strconv.ParseInt(user.ID, 10, 64)

	us, err := models.FindUser(ctx, u.Db, i)

	if err != nil {
		return nil, errors.New("cant find any user for the given id")
	}

	us.Username = user.Username
	us.Email = user.Email

	_, errs := us.Update(ctx, u.Db, boil.Infer())
	if errs != nil {
		return nil, err
	}
	return us, nil

}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, id int64) error {
	rs, err := models.FindUser(ctx, u.Db, id)
	if err != nil {
		return errors.New("cant find any user for the given id")
	}

	_, errs := rs.Delete(ctx, u.Db)
	return errs
}
