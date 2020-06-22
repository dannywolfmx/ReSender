package usecase

import "github.com/dannywolfmx/ReSender/app/domain/model"

//ProfileUseCase represent all the activities a 'profile' can do
type ProfileUseCase interface{

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