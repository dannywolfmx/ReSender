package usecase

import (
	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

//Client
type Client struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type clientUsecase struct {
	repo    app.ClientRepository
	service *service.ClientService
}

//NewClientUsecase retorna una estructura clientUsecase definida por un repositorio y un servidor
func NewClientUsecase(repo app.ClientRepository, service *service.ClientService) *clientUsecase {
	return &clientUsecase{
		repo:    repo,
		service: service,
	}
}

//Clients return a client and error if exist
func (c *clientUsecase) Clients() ([]*Client, error) {
	clientsDomain, err := c.repo.All()

	if err != nil {
		//Return the error
		return nil, err
	}

	clients := make([]*Client, len(clientsDomain))

	//Fill the client with the application model
	for index, client := range clientsDomain {
		clients[index] = &Client{
			ID:   client.ID,
			Name: client.Name,
		}
	}

	return clients, nil

}

//Register add a new client and set the new profile
func (c *clientUsecase) Register(client *Client) error {
	//I don't need the get the ID
	clientDomain := &model.Client{
		Name: client.Name,
	}
	return c.repo.Save(clientDomain)
}

//Delete by id
func (c *clientUsecase) Delete(id uint) error {
	return c.repo.Detele(id)
}

func (c *clientUsecase) Update(client *Client) error {
	domainClient := &model.Client{
		Orm: model.Orm{
			ID: client.ID,
		},
		Name: client.Name,
	}
	return c.repo.Update(domainClient)
}
