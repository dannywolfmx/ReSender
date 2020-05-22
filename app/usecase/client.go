package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

//Un caso de uso de un cliente representa todas las operaciones utiles para un endpoint.
//Una estructura de tipo cliente retornara todas sus estructuras relacionadas

type ClientUsecase interface {
	Clients() ([]*model.Client, error)
	Register(client *model.Client) error
	Delete(id uint) error
	Update(client *model.Client) error
}

func NewClientUsecase(repo repository.Client, service *service.ClientService) *clientUsecase {
	return &clientUsecase{
		repo:    repo,
		service: service,
	}
}

type clientUsecase struct {
	repo    repository.Client
	service *service.ClientService
}

func (c *clientUsecase) Clients() ([]*model.Client, error) {
	return c.repo.All()
}

func (c *clientUsecase) Register(client *model.Client) error {
	return c.repo.Save(client)
}

func (c *clientUsecase) Detele(id uint) error {
	return c.repo.Detele(id)
}

func (c *clientUsecase) Update(client *model.Client) error {
	return c.repo.Update(client)
}
