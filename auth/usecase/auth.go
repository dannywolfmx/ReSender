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

func NewAuthUsecase(
	repo auth.UserRepository,
	service *service.UserService,
) *authUsecase {
	return &authUsecase{
		repo:    repo,
		service: service,
	}
}

func (a *authUsecase) SignUp(username string, password string) error {
	hashPassword, err := a.service.HashAndSaltPassword(password)
	if err != nil {
		return err
	}

	if err := a.service.Duplicated(username); err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Password: hashPassword,
	}

	return a.repo.Create(user)
}

type UserClaims struct {
	jwt.StandardClaims
	User *model.User `json:"user"`
}

func (a *authUsecase) SignIn(username string, password string) (string, error) {
	user, err := a.repo.Get(username)
	if err != nil {
		return "", err
	}

	ok := a.service.ComparePasswordHash(password, user.Password)
	if !ok {

		return "", auth.ErrInvalidPassword
	}

	jwtClaim := UserClaims{
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

	return token.SignedString([]byte("PRUEBA"))
}

func (a *authUsecase) ParseToken(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
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

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidToken
}
