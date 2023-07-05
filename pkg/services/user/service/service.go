package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vier21/go-book-api/config"
	"github.com/vier21/go-book-api/pkg/services/user"
	"github.com/vier21/go-book-api/pkg/services/user/model"
	"github.com/vier21/go-book-api/utils"
)

var (
	ErrEmailAlreadyExist    = errors.New("email is Already Exist")
	ErrUsernameAlreadyExist = errors.New("username is Already Exist")
	ErrInsertUser           = errors.New("failed Insert User")
)

type JWTClaims struct {
	Data utils.LoginPayload `json:"data"`
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

func RegisterResConverter(user model.User) utils.RegisterPayload {
	return utils.RegisterPayload{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (a *User) RegisterUser(ctx context.Context, payload model.User) (utils.RegisterPayload, error) {

	existname, _ := a.UserStore.FindByUsername(ctx, payload.Username)
	existemail, _ := a.UserStore.FindByEmail(ctx, payload.Email)

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

	res, err := a.UserStore.InsertUser(ctx, payload)
	if err != nil {
		return RegisterResConverter(res), ErrInsertUser
	}

	return RegisterResConverter(res), nil
}

func (a *User) LoginUser(ctx context.Context, req utils.LoginRequest) (utils.LoginPayload, string, error) {
	doc, err := a.UserStore.FindByUsername(ctx, req.Username)
	if err != nil {
		return utils.LoginPayload{}, "", errors.New("user not found")
	}

	if err := utils.CheckPasswordHash(req.Password, doc.Password); err != nil {
		return utils.LoginPayload{}, "", errors.New("password not matched")
	}

	mySigningKey := config.GetConfig().SecretKey

	payload := utils.LoginPayload{
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
		payload = utils.LoginPayload{}
		return payload, "", err
	}

	return payload, ss, nil
}
