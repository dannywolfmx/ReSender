package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type OrderUseCase interface {
	DeleteOrder(id uint) error
	UpdateOrder(id uint, number, invoice string) error
}

type orderUsecase struct {
	repo    repository.Order
	service *service.OrderService
}

func NewOrderUsecase(repo repository.Order, service *service.OrderService) *orderUsecase {
	return &orderUsecase{
		repo:    repo,
		service: service,
	}
}

func (o *orderUsecase) DeleteOrder(id uint) error {
	return o.repo.Detele(id)
}
func (o *orderUsecase) UpdateOrder(id uint, number, invoice string) error {
	order := &model.Order{
		Number:  number,
		Invoice: invoice,
	}
	order.ID = id
	return o.repo.Update(order)
}
