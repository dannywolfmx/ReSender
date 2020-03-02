package repository

import (
	"github.com/dannywolfmx/ReSender/models"
	"github.com/rs/xid"
)

//RepositoryClient available methods to manipulate a Client
type (
	Client interface {
		Save(client *models.Client) error
		Get(name string, client *models.Client) error
		Delete(id xid.ID) error
		All(clients *[]models.Client) error
	}

	//RepositoryOrder available methods to manipulate a Order
	Order interface {
		Save(orders *models.Order) error
		Get(id xid.ID, order *models.Order) error
		Delete(id xid.ID) error
		All(order *[]models.Order) error
	}
)
