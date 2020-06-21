package repository

import "github.com/dannywolfmx/ReSender/app/domain/model"

//Order is a repository available methods to implement manipulation on a Order
type Order interface {
	//Delete an order by id
	Detele(id uint) error

	//Update a order
	Update(order *model.Order) error
}

//Client is a repository available methods to implement manipulation on Client model
type Client interface {
	//All get all the clients
	All() ([]*model.Client, error)

	//Delete a client by id
	Detele(id uint) error

	//Find a client by id
	Find(id uint) (*model.Client, error)

	//FindByName find a client by name
	FindByName(name string) (*model.Client, error)

	//Save a client
	Save(client *model.Client) error

	//Update a client
	Update(client *model.Client) error
}

//Mail repositorio para brindar funciones de busqueda, eliminacion, y actualizacion de un correo
//Solo puedo crear un mail por medio de una orden de compra
type Mail interface {

	//Find a mail direction, ej. "test@linux.com"
	Find(direction string) (*model.MailDirection, error)

	//Delete a mail with the id
	Detele(id uint) error

	//DeleteByAddress delete a mail using his mail address
	DeleteByAddress(direction string) error

	//Update a mail direction
	Update(mail *model.MailDirection) error
}

//File repositorio para brindar funciones de busqueda, eliminacion y actualziacion de un los metadatos de un archivo
type File interface {
	//Get a file by id
	Get(id uint) (*model.File, error)

	//Delete a file by id
	Detele(id uint) error

	//Update a file information
	Update(file *model.File) error
}

//Profile repositorio para interactuar con las funcionalidades de una cuenta de usuario
type Profile interface {
	//Get a user account by id
	Get(id uint) (*model.Profile, error)
	//GetByName
	GetByName(name string) (*model.Profile, error)

	//Save a profile
	//Note: the profile name need to be unique
	//**** Check first if the name is already in the database
	Save(profile *model.Profile) error

	//All get all the profiles available
	All() ([]*model.Profile, error)

	//Delete a profile by id
	Detele(id uint) error

	//Update a profile
	Update(profile *model.Profile) error
}
