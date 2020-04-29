package usecase

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

type ClientUseCase interface {
	GetClient(id uint) model.Client
	ListClient() ([]model.Client, error)
	RegisterClient(name string) error
	DeleteClient(id uint) error
	UpdateClient(id uint, name string, orders []model.Order) error
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

func (c *clientUsecase) GetClient(id uint) model.Client {
	return c.repo.GetById(id)
}

func (c *clientUsecase) ListClient() ([]model.Client, error) {
	return c.repo.All()
}

func (c *clientUsecase) RegisterClient(name string) error {
	return c.repo.Save(&model.Client{Name: name})
}

func (c *clientUsecase) DeleteClient(id uint) error {
	return c.repo.Detele(id)
}

func (c *clientUsecase) UpdateClient(id uint, name string, orders []model.Order) error {
	client := &model.Client{
		Name:   name,
		Orders: orders,
	}
	client.ID = id
	return c.repo.Update(client)
}
