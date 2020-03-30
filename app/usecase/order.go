package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type OrderUseCase interface {
	ListOrder() ([]model.Order, error)
	RegisterOrder(number, invoice string, clientid uint) error
	DeleteOrder(invoice string) error
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

func (o *orderUsecase) ListOrder() ([]model.Order, error) {
	return o.repo.All()
}

func (o *orderUsecase) RegisterOrder(number, invoice string, clientid uint) error {
	return o.repo.Save(&model.Order{Number: number, Invoice: invoice, ClientID: clientid})
}
func (o *orderUsecase) DeleteOrder(invoice string) error {
	return o.repo.Detele(invoice)
}
func (o *orderUsecase) UpdateOrder(id uint, number, invoice string) error {
	order := &model.Order{
		Number:  number,
		Invoice: invoice,
	}
	order.ID = id
	return o.repo.Update(order)
}
