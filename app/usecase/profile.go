package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

//ProfileUsecase represent all the activities a 'profile' can do
type ProfileUsecase interface {
	//Create a new profile and return a nil error if the transactions workds.
	GetAll() ([]*model.Profile, error)

	//Create a new profile and return a nil error if the transactions workds.
	GetByID(profileID uint) (*model.Profile, error)

	//Create a new profile and return a nil error if the transactions workds.
	Create(profile *model.Profile) error

	//Create password to the profile and return an error if the transaction doesnt work
	UpdatePassword(profileID uint, password string) error

	//Add a new client to the profile client list
	//Search a profile by ID
	//Set a relationship beetween the new client
	AddClient(profileID uint, client *model.Client) error

	//Delete profile account
	Delete(profileID uint) error

	//Update a profile, return the new profile and error
	Update(profile *model.Profile) error
}

type profileUsecase struct {
	repo    repository.Profile
	service *service.ProfileService
}

//NewProfileUsecase create a new profile structure with the repository and the service
func NewProfileUsecase(repo repository.Profile, service *service.ProfileService) *profileUsecase {
	return &profileUsecase{
		repo:    repo,
		service: service,
	}
}

//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) GetAll() ([]*model.Profile, error) {
	return u.repo.All()
}

//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) GetByID(profileID uint) (*model.Profile, error) {
	return u.repo.Get(profileID)
}

//TODO implment ProfileUsecase de profileUsecase
//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) Create(profile *model.Profile) error {

	//Check if the name is already in the data base
	err := u.service.Duplicated(profile.Name)
	if err != nil {
		return err
	}

	//Hash the password
	hash, err := u.service.HashAndSaltPassword(profile.Password)
	if err != nil {
		return err
	}

	profile.Password = hash

	//Save the profile an check errors
	err = u.repo.Save(profile)
	if err != nil {
		return err
	}

	return nil
}

//Create password to the profile and return an error if the transaction doesnt work
func (u *profileUsecase) UpdatePassword(profileID uint, password string) error {
	//Get the profile by id
	profile, err := u.repo.Get(profileID)
	if err != nil {
		return err
	}

	//Hash the password
	hash, err := u.service.HashAndSaltPassword(password)
	if err != nil {
		return err
	}

	profile.Password = hash

	//Update password
	err = u.repo.Update(profile)

	if err != nil {
		return err
	}

	return nil
}

//Add a new client to the profile client list
//Search a profile by ID
//Set a relationship beetween the new client
func (u *profileUsecase) AddClient(profileID uint, client *model.Client) error {
	//Get the profile by id
	profile, err := u.repo.Get(profileID)
	if err != nil {
		return err
	}

	//Append the profile
	profile.Clients = append(profile.Clients, client)

	//try to update and return a error if exist
	return u.repo.Update(profile)

}

//Delete profile account
func (u *profileUsecase) Delete(profileID uint) error {
	return u.repo.Detele(profileID)
}

//Update a profile, return the new profile and error
func (u *profileUsecase) Update(profile *model.Profile) error {
	return u.repo.Update(profile)
}
