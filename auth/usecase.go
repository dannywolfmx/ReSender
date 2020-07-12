package auth

import "github.com/dannywolfmx/ReSender/auth/domain/model"

const ContexUserKey = "user"

type AuthUsecase interface {
	//SignUp return a token and error
	SignUp(username, password string) (string, error)
	SignIn(username, password string) (string, error)
	ParseToken(token string) (*model.User, error)
}
