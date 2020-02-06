package service

import (
	"github.com/dannywolfmx/ReSender/models"
	"github.com/rs/xid"
)

//Order ServiceOrder
type Order interface {
	All(orders *[]models.Order) error
	Delete(id xid.ID) error
	Get(id xid.ID, order *models.Order) error
	Save(order *models.Order) error
}
