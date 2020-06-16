package repository

import "github.com/dannywolfmx/ReSender/app/domain/model"

//Order is a repository available methods to implement manipulation on a Order
type Order interface {
	Detele(id uint) error
	Update(*model.Order) error
}

//Client is a repository available methods to implement manipulation on Client model
type Client interface {
	Save(*model.Client) error
	FindByName(string) (*model.Client, error)
	GetById(id uint) model.Client
	All() ([]*model.Client, error)
	Detele(id uint) error
	Update(*model.Client) error
}

//Mail repositorio para brindar funciones de busqueda, eliminacion, y actualizacion de un correo
//Solo puedo crear un mail por medio de una orden de compra
type Mail interface {
	Find(string) (*model.MailDirection, error)
	Detele(id uint) error
	Update(*model.MailDirection) error
}

//File repositorio para brindar funciones de busqueda, eliminacion y actualziacion de un los metadatos de un archivo
type File interface {
	Get(id uint) (*model.File, error)
	Detele(id uint) error
	Update(*model.File) error
}

//Account repositorio para interactuar con las funcionalidades de una cuenta de usuario
type Account interface {
	Get(id uint) (*model.Profile, error)
	Save(*model.Profile) error
	All() ([]*model.Profile, error)
	Detele(id uint) error
	Update(*model.Profile) error
}
