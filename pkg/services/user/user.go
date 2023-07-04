package user

import (
	"context"

	"github.com/vier21/go-book-api/pkg/services/user/model"
	"github.com/vier21/go-book-api/utils"
)

type UserRepo interface {
	FindByUsername(context.Context, string) (model.User, error)
	FindById(context.Context, string) (model.User, error)
	InsertUser(context.Context, model.User) (model.User, error)
	UpdateUser(context.Context, model.User) (model.User, error)
	DeleteUser(context.Context, ...string) error
}

type UserService interface {
	RegisterUser(context.Context, model.User) (model.User, error)
	LoginUser(context.Context, utils.LoginRequest) (string, error)
}
