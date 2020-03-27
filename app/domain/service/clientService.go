package service

import (
	"fmt"

	"github.com/dannywolfmx/ReSender/app/domain/repository"
)

//Nota ClientService debe ser publico, dado que el usecase hace uso de este tipo
type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{
		repo: repo,
	}
}

//Duplicated check if a name is already created
func (s *ClientService) Duplicated(name string) error {
	client, err := s.repo.FindByName(name)
	if client != nil {
		return fmt.Errorf("%s already exists", name)
	}

	return err
}
