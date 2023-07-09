package user

import (
	"context"

	"github.com/vier21/go-book-api/pkg/services/user/def"
	"github.com/vier21/go-book-api/pkg/services/user/model"
)

type UserRepo interface {
	FindByUsername(context.Context, string) (model.User, error)
	FindByEmail(context.Context, string) (model.User, error)
	FindById(context.Context, string) (model.User, error)
	InsertUser(context.Context, model.User) (model.User, error)
	UpdateUser(context.Context, string, model.UpdateUser) (model.User, error)
	DeleteUser(context.Context, string) error
	BulkDelete(context.Context, ...string) error
}

type UserService interface {
	RegisterUser(context.Context, model.User) (def.RegisterPayload, error)
	LoginUser(context.Context, def.LoginRequest) (def.LoginPayload, string, error)
	UpdateUser(context.Context, string, model.UpdateUser) (def.UpdatePayload, error)
	DeleteUser(context.Context, string) error
	BulkDeleteUser(context.Context, []string) error
}
