package repository

import "github.com/dannywolfmx/ReSender/app/domain/model"

//Order is a repository available methods to manipulate a Order
type Order interface {
	Save(*model.Order) error
	FindByInvoice(string) (*model.Order, error)
	All() ([]model.Order, error)
	Detele(invoice string) error
	Update(*model.Order) error
}
