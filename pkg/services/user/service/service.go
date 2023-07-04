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

func (a *User) RegisterUser(ctx context.Context, payload model.User) (model.User, error) {

	exist, _ := a.UserStore.FindByUsername(ctx, payload.Username)

	if exist.Username == payload.Username {
		if exist.Email == payload.Email {
			return exist, ErrEmailAlreadyExist
		}
		return exist, ErrUsernameAlreadyExist
	}

	pass, err := utils.HashPassword(payload.Password)

	payload.Password = pass
	if err != nil {
		return model.User{}, err
	}

	res, err := a.UserStore.InsertUser(ctx, payload)
	if err != nil {
		return res, ErrInsertUser
	}

	return res, nil
}

func (a *User) LoginUser(ctx context.Context, req utils.LoginRequest) (string, error) {
	doc, err := a.UserStore.FindByUsername(ctx, req.Username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := utils.CheckPasswordHash(req.Password, doc.Password); err != nil {
		return "", errors.New("password not matched")
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
		return "", err
	}

	return ss, nil
}
