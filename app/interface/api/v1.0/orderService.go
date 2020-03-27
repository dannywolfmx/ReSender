package v1

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/usecase"
)

type orderService struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderService(orderUseCase usecase.OrderUseCase) *orderService {
	return &orderService{
		orderUseCase: orderUseCase,
	}
}

func (s *orderService) ListOrder() ([]model.Order, error) {
	return s.orderUseCase.ListOrder()
}

func (s *orderService) RegisterOrder(number, invoice string) error {
	return s.orderUseCase.RegisterOrder(number, invoice)
}

func (s *orderService) DeleteOrder(invoice string) error {
	return s.orderUseCase.DeleteOrder(invoice)
}

func (s *orderService) UpdateOrder(id uint, number, invoice string) error {
	return s.orderUseCase.UpdateOrder(id, number, invoice)
}
