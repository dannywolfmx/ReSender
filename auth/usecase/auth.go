package usecase

import (
	"fmt"
	"time"

	"github.com/dannywolfmx/ReSender/auth"
	"github.com/dannywolfmx/ReSender/auth/domain/model"
	"github.com/dannywolfmx/ReSender/auth/domain/service"
	"github.com/dgrijalva/jwt-go"
)

type authUsecase struct {
	repo           auth.UserRepository
	service        *service.UserService
	hashSalt       string
	signKey        []byte
	expireDuration time.Duration
}

func NewAuthUsecase(repo auth.UserRepository, service *service.UserService) *authUsecase {
	return &authUsecase{
		repo:    repo,
		service: service,
	}
}

func (a *authUsecase) SignUp(username string, password string) (string, error) {
	hashPassword, err := a.service.HashAndSaltPassword(password)
	if err != nil {
		return "", err
	}

	//The username already exist
	if err := a.service.Duplicated(username); err != nil {
		return "", auth.ErrNameAlreayExist
	}

	user := &model.User{
		Username: username,
		Password: hashPassword,
	}
	err = a.repo.Create(user)

	if err != nil {
		return "", err
	}

	return a.service.GetJWTToken(user)
}

func (a *authUsecase) SignIn(username string, password string) (string, error) {
	user, err := a.repo.Get(username)
	if err != nil {
		return "", err
	}

	//The user doesnt exist
	if user == nil {
		return "", auth.ErrInvalidUser
	}

	ok := a.service.ComparePasswordHash(password, user.Password)
	if !ok {

		return "", auth.ErrInvalidPassword
	}

	return a.service.GetJWTToken(user)
}

func (a *authUsecase) ParseToken(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("PRUEBA"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
		user, err := a.repo.Get(claims.User.Username)
		if err != nil || user == nil {
			return nil, auth.ErrInvalidToken
		}
		return user, nil
	}

	return nil, auth.ErrInvalidToken
}
