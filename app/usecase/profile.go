package usecase

import (
	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type profileUsecase struct {
	repo    app.ProfileRepository
	service *service.ProfileService
}

//NewProfileUsecase create a new profile structure with the repository and the service
func NewProfileUsecase(repo app.ProfileRepository, service *service.ProfileService) *profileUsecase {
	return &profileUsecase{
		repo:    repo,
		service: service,
	}
}

//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) GetAll() ([]*model.Profile, error) {
	//Return all the profiles
	return u.repo.All()
}

//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) GetByID(profileID uint) (*model.Profile, error) {
	//Get just one result
	return u.repo.Get(profileID)

}

//TODO implment ProfileUsecase de profileUsecase
//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) Create(imageAvatarPath, name, password string) error {

	//Check if the name is already in the data base
	err := u.service.Duplicated(name)
	if err != nil {
		return err
	}

	//Hash the password
	hash, err := u.service.HashAndSaltPassword(password)
	if err != nil {
		return err
	}

	profile := &model.Profile{
		ImageAvatarPath: imageAvatarPath,
		Name:            name,
		//Important hash the password first
		Password: hash,
	}

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

//Delete profile account
func (u *profileUsecase) Delete(profileID uint) error {
	return u.repo.Detele(profileID)
}

//Update a profile, return the new profile and error
func (u *profileUsecase) Update(profileID uint, imageAvatarPath, name string) error {

	//Transform the data to domain entity
	profile := &model.Profile{
		ImageAvatarPath: imageAvatarPath,
		Name:            name,
	}

	//Send the update value
	return u.repo.Update(profile)
}
