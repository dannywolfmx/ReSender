//IMPLEMENTACION DEL SERVICIO CLIENTE
package v1

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/usecase"
)

type clientService struct {
	clientUseCase usecase.ClientUseCase
}

func NewClientService(clientUseCase usecase.ClientUseCase) *clientService {
	return &clientService{
		clientUseCase: clientUseCase,
	}
}

func (c *clientService) GetClient(id uint) model.Client {
	return c.clientUseCase.GetClient(id)
}

func (c *clientService) ListClient() ([]model.Client, error) {
	return c.clientUseCase.ListClient()
}

func (c *clientService) RegisterClient(name string) error {
	return c.clientUseCase.RegisterClient(name)
}

func (c *clientService) DeleteClient(id uint) error {
	return c.clientUseCase.DeleteClient(id)
}

func (c *clientService) UpdateClient(id uint, name string, orders []model.Order) error {
	return c.clientUseCase.UpdateClient(id, name, orders)
}
