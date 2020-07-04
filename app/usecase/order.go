package usecase

import (
	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

//OrderUsecase define un usecase para el modelo Order
type OrderUsecase interface {
	Delete(id uint) error
	Update(id uint, number, invoice string) error
}

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

func (o *orderUsecase) Delete(id uint) error {
	return o.repo.Detele(id)
}
func (o *orderUsecase) Update(id uint, number, invoice string) error {
	order := &model.Order{
		Number:  number,
		Invoice: invoice,
	}
	order.ID = id
	return o.repo.Update(order)
}
