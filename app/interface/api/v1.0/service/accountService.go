package service

import "github.com/dannywolfmx/ReSender/app/usecase"

type accountService struct {
	accountUsecase usecase.AccountUsecase
}

func NewAccountService(accountUsecase usecase.AccountUsecase) *accountService {
	return &accountService{
		accountUsecase: accountUsecase,
	}
}
