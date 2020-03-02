package logic

import (
	"time"

	"github.com/dannywolfmx/ReSender/models"
	"github.com/dannywolfmx/ReSender/models/service"
	"github.com/dannywolfmx/ReSender/repository"
	"github.com/rs/xid"
)

type orderService struct {
	repo repository.Order
}

func NewOrderService(orderRepo repository.Order) service.Order {
	return &orderService{
		orderRepo,
	}
}

func (o *orderService) All(orders *[]models.Order) error {
	return o.repo.All(orders)
}

func (o *orderService) Delete(id xid.ID) error {
	return o.repo.Delete(id)
}

func (o *orderService) Get(id xid.ID, order *models.Order) error {
	return o.repo.Get(id, order)
}

func (o *orderService) Save(order *models.Order) error {
	//Realizar validaciones antes de guardar
	order.ID = xid.New()
	order.CreatedAt = time.Now().UTC().Unix()
	return o.repo.Save(order)
}
