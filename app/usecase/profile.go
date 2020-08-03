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

//create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) GetByUserID(id uint) (*model.Profile, error) {
	//get just one result
	return u.repo.GetByUserID(id)

}

//TODO implment ProfileUsecase de profileUsecase
//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) Create(userID uint) (*model.Profile, error) {

	profile := &model.Profile{
		UserID: userID,
	}

	//Save the profile an check errors
	err := u.repo.Save(profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

//Delete profile account
func (u *profileUsecase) Delete(profileID uint) error {
	return u.repo.Detele(profileID)
}

//Update a profile, return the new profile and error
func (u *profileUsecase) Update(profileID uint, imageAvatarPath string) error {

	//Transform the data to domain entity
	profile := &model.Profile{
		ImageAvatarPath: imageAvatarPath,
	}

	//Send the update value
	return u.repo.Update(profile)
}
