package auth

import "github.com/dannywolfmx/ReSender/auth/domain/model"

type AuthUsecase interface {
	SignUp(username, password string) error
	SignIn(username, password string) error
	ParseToken(token string) (*model.User, error)
}
