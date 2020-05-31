package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type AccountUsecase interface {
	Get(id uint) (*model.Account, error)
	List() ([]*model.Account, error)
	Delete(id uint) error
	Update(account *model.Account) error
	RegisterUser(account *model.Account) error
}

type accountUsecase struct {
	repo    repository.Account
	service *service.AccountService
}

func (a *accountUsecase) Get(id uint) (*model.Account, error) {
	return a.repo.Get(id)
}

func (a *accountUsecase) List() ([]*model.Account, error) {
	return a.repo.All()
}

func (a *accountUsecase) Delete(id uint) error {
	return a.repo.Detele(id)
}

func (a *accountUsecase) Update(account *model.Account) error {
	return a.repo.Update(account)
}

func (a *accountUsecase) RegisterUser(account *model.Account) error {
	return a.repo.Save(account)
}