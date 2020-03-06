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

func (s *orderService) ListOrder() ([]*model.Order, error) {
	orders, err := s.orderUseCase.ListOrder()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderService) RegisterOrder(number, invoice string) error {
	err := s.orderUseCase.RegisterOrder(number, invoice)
	if err != nil {
		return err
	}
	return nil
}
