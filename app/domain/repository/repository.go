package repository

import "github.com/dannywolfmx/ReSender/app/domain/model"

//Order is a repository available methods to implement manipulation on a Order
type Order interface {
	Save(*model.Order) error
	FindByInvoice(invoice string) (*model.Order, error)
	GetById(id uint) model.Order
	All() ([]model.Order, error)
	Detele(id uint) error
	Update(*model.Order) error
}

//Client is a repository available methods to implement manipulation on Client model
type Client interface {
	Save(*model.Client) error
	FindByName(string) (*model.Client, error)
	GetById(id uint) model.Client
	//TODO: Convertir esta lista a una lista de apuntadores "[]*model.Client"
	All() ([]model.Client, error)
	Detele(id uint) error
	Update(*model.Client) error
}
