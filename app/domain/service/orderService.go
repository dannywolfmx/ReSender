package service

import (
	"github.com/dannywolfmx/ReSender/app/domain/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}
