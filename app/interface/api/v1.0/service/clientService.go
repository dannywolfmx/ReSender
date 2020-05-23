//IMPLEMENTACION DEL SERVICIO CLIENTE
package service

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/usecase"
)

type clientService struct {
	clientUseCase usecase.ClientUsecase
}

func NewClientService(clientUseCase usecase.ClientUsecase) *clientService {
	return &clientService{
		clientUseCase: clientUseCase,
	}
}

func (c *clientService) ListClient() ([]*model.Client, error) {
	return c.clientUseCase.Clients()
}

func (c *clientService) RegisterClient(client *model.Client) error {
	return c.clientUseCase.Register(client)
}

func (c *clientService) DeleteClient(id uint) error {
	return c.clientUseCase.Delete(id)
}

func (c *clientService) UpdateClient(client *model.Client) error {
	return c.clientUseCase.Update(client)
}
