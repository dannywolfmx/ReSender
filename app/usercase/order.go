package usercase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type OrderUseCase interface {
	ListOrder() ([]*model.Order, error)
	RegisterOrder(string, string) error
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

func (o *orderUsecase) ListOrder() ([]*model.Order, error) {
	orders, err := o.repo.All()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderUsecase) RegisterOrder(number, invoice string) error {
	return nil
}
