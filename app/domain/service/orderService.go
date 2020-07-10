package service

import "github.com/dannywolfmx/ReSender/app"

type OrderService struct {
	repo app.OrderRepository
}

func NewOrderService(repo app.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}
