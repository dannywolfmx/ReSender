package service

import (
	"fmt"

	"github.com/dannywolfmx/ReSender/app/domain/repository"
)

type OrderService struct {
	repo repository.Order
}

func (s *OrderService) Duplicated(invoice string) error {
	order, err := s.repo.FindByInvoice(invoice)
	if order != nil {
		return fmt.Errorf("%s already exists", invoice)
	}

	if err != nil {
		return err
	}
	return nil
}
