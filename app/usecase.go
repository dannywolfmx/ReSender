package app

import "github.com/dannywolfmx/ReSender/app/domain/model"

//Put here all your usecase from the app

//ClientUsecase Un caso de uso de un cliente representa todas las operaciones utiles para un endpoint.
//Una estructura de tipo cliente retornara todas sus estructuras relacionadas
type ClientUsecase interface {
	//Clients list of clients
	Clients() ([]*model.Client, error)

	//Register add a new client and set the new profile
	Register(profileID uint, name string) error

	//Delete a client by id
	Delete(id uint) error

	//Update a client
	Update(id uint, name string) error
}

//ProfileUsecase represent all the activities a 'profile' can do
type ProfileUsecase interface {
	//Create a new profile and return a nil error if the transactions workds.
	GetAll() ([]*model.Profile, error)

	//Create a new profile and return a nil error if the transactions workds.
	GetByID(profileID uint) (*model.Profile, error)

	//Create a new profile and return a nil error if the transactions workds.
	Create(imageAvatarPath, name, password string) error

	//Delete profile account
	Delete(profileID uint) error

	//Update a profile, return the new profile and error
	Update(profileID uint, imageAvatarPath, name string) error

	//Create password to the profile and return an error if the transaction doesnt work
	UpdatePassword(profileID uint, password string) error
}

//OrderUsecase define un usecase para el modelo Order
type OrderUsecase interface {
	Delete(id uint) error
	Update(id uint, number, invoice string) error
}
