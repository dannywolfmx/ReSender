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
	Update(profile *model.Profile) (*model.Profile, error)
}

type profileUsecase struct {
	repo    repository.Profile
	service *service.ProfileService
}

func NewProfileUsecase(repo repository.Profile, service *service.ProfileService) *profileUsecase {
	return &profileUsecase{
		repo:    repo,
		service: service,
	}
}

//TODO implment ProfileUsecase de profileUsecase
//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) Create(profile *model.Profile) error {

	//Check if the name is already in the data base
	err := u.service.Duplicated(profile.Name)
	if err != nil {
		return err
	}

	//Save the profile an check errors
	err = u.repo.Save(profile)
	if err != nil {
		return err
	}

	return nil
}

//Create password to the profile and return an error if the transaction doesnt work
func (u *profileUsecase) SetPassword(password string) error {
	panic("not implemented") // TODO: Implement
}

//Add a new client to the profile client list
//Search a profile by ID
//Set a relationship beetween the new client
func (u *profileUsecase) AddClient(profileID uint, client *model.Client) error {
	panic("not implemented") // TODO: Implement
}

//Delete profile account
func (u *profileUsecase) Delete(profileID uint) (*model.Profile, error) {
	panic("not implemented") // TODO: Implement
}

//Update a profile, return the new profile and error
func (u *profileUsecase) Update(profile *model.Profile) (*model.Profile, error) {
	panic("not implemented") // TODO: Implement
}
