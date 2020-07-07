package usecase

import (
	"github.com/dannywolfmx/ReSender/auth"
	"github.com/dannywolfmx/ReSender/auth/domain/model"
	"github.com/dannywolfmx/ReSender/auth/domain/service"
)

type authUsecase struct {
	repo    auth.UserRepository
	service *service.UserService
}

func (a *authUsecase) SignUp(username string, password string) error {
	panic("not implemented") // TODO: Implement
}

func (a *authUsecase) SignIn(username string, password string) error {
	panic("not implemented") // TODO: Implement
}

func (a *authUsecase) ParseToken(token string) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func NewAuthUsecase(repo auth.UserRepository, service *service.UserService) *authUsecase {
	return &authUsecase{
		repo:    repo,
		service: service,
	}
}
