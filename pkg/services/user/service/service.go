package service

import (
	"context"
	"errors"

	"github.com/vier21/go-book-api/pkg/services/user"
	"github.com/vier21/go-book-api/pkg/services/user/model"
	"github.com/vier21/go-book-api/utils"
)

var ErrEmailAlreadyExist = errors.New("Email is Already Exist")
var ErrUsernameAlreadyExist = errors.New("Username is Already Exist")
var ErrInsertUser = errors.New("Failed Insert User")

type Auth struct {
	UserStore user.UserRepo
}

func NewUserService(auth user.UserRepo) *Auth {
	return &Auth{
		UserStore: auth,
	}
}

func (a *Auth) RegisterUser(ctx context.Context, payload model.User) (model.User, error) {

	exist, err := a.UserStore.FindByUsername(ctx, payload.Username)

	if err != nil {
		return model.User{}, err
	}

	if exist.Username == payload.Username {
		if exist.Email == payload.Email {
			return exist, ErrEmailAlreadyExist
		}
		return exist, ErrUsernameAlreadyExist
	}

	payload.Password, err = utils.HashPassword(payload.Password)

	if err != nil {
		return model.User{}, err
	}

	res, err := a.UserStore.InsertUser(ctx, payload)

	if err != nil {
		return res, ErrInsertUser
	}

	return res, nil
}

func (a *Auth) LoginUser(ctx context.Context, payload utils.LoginRequest) (utils.LoginPayload, error) {

	return utils.LoginPayload{}, nil
}
