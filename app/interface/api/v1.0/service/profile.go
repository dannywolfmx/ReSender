package service

import "github.com/dannywolfmx/ReSender/app/usecase"

type profileService struct {
	u usecase.ProfileUsecase
}

func NewProficeService(u usecase.ProfileUsecase) *profileService {
	return &profileService{
		u: u,
	}
}
