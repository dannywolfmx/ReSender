package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type ClientUseCase interface {
	ListClient() ([]model.Client, error)
	RegisterClient(name string) error
	DeleteClient(name string) error
	UpdateClient(id uint, name string) error
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

func (c *clientUsecase) ListClient() ([]model.Client, error) {
	return c.repo.All()
}

func (c *clientUsecase) RegisterClient(name string) error {
	return c.repo.Save(&model.Client{Name: name})
}

func (c *clientUsecase) DeleteClient(name string) error {
	return c.repo.Detele(name)
}

func (c *clientUsecase) UpdateClient(id uint, name string) error {
	client := &model.Client{
		Name: name,
	}
	client.ID = id
	return c.repo.Update(client)
}
