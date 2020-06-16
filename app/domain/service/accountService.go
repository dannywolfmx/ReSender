package service

import (
	"github.com/dannywolfmx/ReSender/app/domain/repository"
)

//Nota ClientService debe ser publico, dado que el usecase hace uso de este tipo
type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

//Hash password
func (a *AccountService) hashPassword(pass string) string {
	panic("implementar ")
}
