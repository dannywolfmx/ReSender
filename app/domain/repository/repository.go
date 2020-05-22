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
	//TODO: Convertir esta lista a una lista de apuntadores "[]*model.Client"
	All() ([]*model.Client, error)
	Detele(id uint) error
	Update(*model.Client) error
}

//Solo puedo crear un mail por medio de una orden de compra
type Mail interface {
	Find(string) (*model.MailDirection, error)
	Detele(id uint) error
	Update(*model.MailDirection) error
}

type File interface {
	Get(id uint) (*model.File, error)
	Detele(id uint) error
	Update(*model.File) error
}

type Account interface {
	Save(*model.Account) error
	All() ([]*model.Account, error)
	Detele(id uint) error
	Update(*model.Account) error
}
