package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type AccountUsecase interface {
	Get(id uint) (*model.Profile, error)
	List() ([]*model.Profile, error)
	Delete(id uint) error
	Update(account *model.Profile) error
	RegisterUser(account *model.Profile) error
}

type accountUsecase struct {
	repo    repository.Account
	service *service.AccountService
}

func NewAccountUseCase(repo repository.Account, service *service.AccountService) *accountUsecase {
	return &accountUsecase{
		repo:    repo,
		service: service,
	}
}

func (a *accountUsecase) Get(id uint) (*model.Profile, error) {
	return a.repo.Get(id)
}

func (a *accountUsecase) List() ([]*model.Profile, error) {
	return a.repo.All()
}

func (a *accountUsecase) Delete(id uint) error {
	return a.repo.Detele(id)
}

func (a *accountUsecase) Update(account *model.Profile) error {
	return a.repo.Update(account)
}

func (a *accountUsecase) RegisterUser(account *model.Profile) error {
	return a.repo.Save(account)
}
