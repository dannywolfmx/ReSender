package repository

import "github.com/dannywolfmx/ReSender/app/domain/model"

//RepositoryOrder available methods to manipulate a Order
type Order interface {
	Save(*model.Order) error
	FindByInvoice(string) (*model.Order, error)
	All() ([]*model.Order, error)
}
