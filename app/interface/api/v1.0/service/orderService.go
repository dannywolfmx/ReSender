//IMPLEMENTACION DEL SERVICIO ORDENES
package service

import (
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

func (s *orderService) DeleteOrder(id uint) error {
	return s.orderUseCase.DeleteOrder(id)
}

func (s *orderService) UpdateOrder(id uint, number, invoice string) error {
	return s.orderUseCase.UpdateOrder(id, number, invoice)
}
