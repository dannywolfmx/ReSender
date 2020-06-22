package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

//ProfileUseCase represent all the activities a 'profile' can do
type ProfileUsecase interface {

	//Create a new profile and return a nil error if the transactions workds.
	Create(profile *model.Profile) error

	//Create password to the profile and return an error if the transaction doesnt work
	SetPassword(password string) error

	//Add a new client to the profile client list
	//Search a profile by ID
	//Set a relationship beetween the new client
	AddClient(profileID uint, client *model.Client) error

	//Delete profile account
	Delete(profileID uint) (*model.Profile, error)

	//Update a profile, return the new profile and error
	Update(profile *mode.Profile) (*model.Profile, error)
}

type profileUsecase struct {
	repo    repository.Profile
	service *service.ProfileService
}

func NewProfileUsecase(repo repository.Profile, service *service.ProfileService) *profileprofileUsecase {
	return &profileUsecase{
		reto:    repo,
		service: service,
	}
}

//TODO implment ProfileUsecase de profileUsecase
