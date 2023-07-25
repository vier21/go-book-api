package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vier21/go-book-api/config"
	"github.com/vier21/go-book-api/pkg/services/user"
	"github.com/vier21/go-book-api/pkg/services/user/common"
	"github.com/vier21/go-book-api/pkg/services/user/model"
	"github.com/vier21/go-book-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrEmailAlreadyExist    = errors.New("email is Already Exist")
	ErrUsernameAlreadyExist = errors.New("username is Already Exist")
	ErrInsertUser           = errors.New("failed Insert User")
)

type JWTClaims struct {
	Data common.LoginPayload `json:"data"`
	jwt.RegisteredClaims
}

type User struct {
	UserStore user.UserRepo
}

func NewUserService(auth user.UserRepo) *User {
	return &User{
		UserStore: auth,
	}
}

func RegisterResConverter(usr model.User) common.RegisterPayload {
	return common.RegisterPayload{
		Id:       usr.Id,
		Username: usr.Username,
		Email:    usr.Email,
	}
}

func (u *User) RegisterUser(ctx context.Context, payload model.User) (common.RegisterPayload, error) {

	existname, _ := u.UserStore.FindByUsername(ctx, payload.Username)
	existemail, _ := u.UserStore.FindByEmail(ctx, payload.Email)

	email := existemail.Email
	username := existname.Username

	if username == payload.Username {
		return RegisterResConverter(existname), ErrUsernameAlreadyExist
	}

	if email == payload.Email {
		return RegisterResConverter(existemail), ErrEmailAlreadyExist
	}

	pass, err := utils.HashPassword(payload.Password)

	payload.Password = pass
	if err != nil {
		return RegisterResConverter(model.User{}), err
	}

	res, err := u.UserStore.InsertUser(ctx, payload)
	if err != nil {
		return RegisterResConverter(res), ErrInsertUser
	}

	return RegisterResConverter(res), nil
}

func (u *User) LoginUser(ctx context.Context, req common.LoginRequest) (common.LoginPayload, string, error) {
	doc, err := u.UserStore.FindByUsername(ctx, req.Username)
	if err != nil {
		return common.LoginPayload{}, "", errors.New("user not found")
	}

	if err := utils.CheckPasswordHash(req.Password, doc.Password); err != nil {
		return common.LoginPayload{}, "", errors.New("password not matched")
	}

	mySigningKey := config.GetConfig().SecretKey

	payload := common.LoginPayload{
		Id:       doc.Id,
		Username: doc.Username,
		Email:    doc.Email,
	}

	claims := JWTClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "www.xavdoc.me",
			Subject:   doc.Id,
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		payload = common.LoginPayload{}
		return payload, "", err
	}

	return payload, ss, nil
}

func (u *User) UpdateUser(ctx context.Context, id string, payload model.UpdateUser) (common.UpdatePayload, error) {
	empty := model.UpdateUser{}
	_, err := u.UserStore.FindById(ctx, id)

	if id == "" {
		return common.UpdatePayload{}, errors.New("no is specified is specified in parameter")
	}

	if err == mongo.ErrNoDocuments {
		return common.UpdatePayload{}, errors.New("no matched user with given id")
	}

	if payload == empty {
		return common.UpdatePayload{}, errors.New("your data is up to date")
	}
	
	doc, err := u.UserStore.UpdateUser(ctx, id, payload)
	if err != nil {
		return common.UpdatePayload{}, err
	}

	result := updatedPayload(doc)
	return result, nil
}

func (u *User) DeleteUser(ctx context.Context, id string) error {
	err := u.UserStore.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BulkDeleteUser(ctx context.Context, bulk []string) error {
	err := u.UserStore.BulkDelete(ctx, bulk...)
	if err != nil {
		return err
	}
	return nil
}

func updatedPayload(upUser model.User) common.UpdatePayload {
	return common.UpdatePayload{
		Id:       upUser.Id,
		Username: upUser.Username,
		Email:    upUser.Email,
	}
}
