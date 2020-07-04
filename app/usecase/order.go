package usecase

import (
	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type orderUsecase struct {
	repo    app.OrderRepository
	service *service.OrderService
}

//NewOrderUsecase construlle una orderUsecase bien definido
func NewOrderUsecase(repo app.OrderRepository, service *service.OrderService) *orderUsecase {
	return &orderUsecase{
		repo:    repo,
		service: service,
	}
}

//Delet e a order
func (o *orderUsecase) Delete(id uint) error {
	return o.repo.Detele(id)
}

//Update a order with a valid id
func (o *orderUsecase) Update(id uint, number, invoice string) error {
	order := &model.Order{
		Number:  number,
		Invoice: invoice,
	}
	order.ID = id
	return o.repo.Update(order)
}
