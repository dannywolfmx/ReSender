package usecase

import (
	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

//Profile
type Profile struct {
	//ImageAvatarPath Imagen del perfil
	ImageAvatarPath string `json:"image_avatar_path"`

	//Name nombre del perfil
	Name string `json:"name"`
}

//ProfileWithPassowrd use this structure to unmarshal the profile with the password
type ProfileWithPassowrd struct {
	//ImageAvatarPath Imagen del perfil
	ImageAvatarPath string `json:"image_avatar_path"`

	//Name nombre del perfil
	Name     string `json:"name"`
	Password string `json:"password"`
}

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
func (u *profileUsecase) GetAll() ([]*Profile, error) {
	profilesDomain, err := u.repo.All()

	profiles := make([]*Profile, len(profilesDomain))

	for index, profile := range profilesDomain {
		profiles[index] = &Profile{
			ImageAvatarPath: profile.ImageAvatarPath,
			Name:            profile.Name,
		}
	}

	return profiles, err
}

//Create a new profile and return a nil error if the transactions workds.
func (u *profileUsecase) GetByID(profileID uint) (*Profile, error) {
	profileDomain, err := u.repo.Get(profileID)

	profile := &Profile{
		ImageAvatarPath: profileDomain.ImageAvatarPath,
		Name:            profileDomain.Name,
	}

	return profile, err
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
func (u *profileUsecase) Update(profile *Profile) error {

	//Transform the data to domain entity
	profileDomain := &model.Profile{
		ImageAvatarPath: profile.ImageAvatarPath,
		Name:            profile.Name,
	}

	//Send the update value
	return u.repo.Update(profileDomain)
}

//Update a profile, return the new profile and error
func (u *profileUsecase) UpdateWithPassword(profile *ProfileWithPassowrd) error {
	//Transform the data to domain entity
	profileDomain := &model.Profile{
		ImageAvatarPath: profile.ImageAvatarPath,
		Name:            profile.Name,
		Password:        profile.Password,
	}

	//Send the update value
	return u.repo.Update(profileDomain)
}
