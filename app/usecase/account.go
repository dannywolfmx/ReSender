package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type AccountUseCase interface {
	Get(id uint) model.Account
	List() ([]model.Account, error)
	Save() error
	Delete(id uint) error
	Update(id uint) error
}

type accountUseCase struct {
	repo    repository.Account
	service *service.AccountService
}

func NewAccountUseCase(r repository.Account, s *service.AccountService) *accountUseCase {
	return &accountUseCase{
		repo:    r,
		service: s,
	}
}
