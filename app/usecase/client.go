package usecase

import (
	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
)

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
func (c *clientUsecase) Clients() ([]*model.Client, error) {
	return c.repo.All()
}

//Register add a new client and set the new profile
func (c *clientUsecase) Register(profileID uint, name string) error {
	//I don't need the get the ID
	//Conver the data to model.Client struct
	client := &model.Client{
		ProfiletID: profileID,
		Name:       name,
	}
	return c.repo.Save(client)
}

//Delete by id
func (c *clientUsecase) Delete(id uint) error {
	return c.repo.Detele(id)
}

//Update the client name
func (c *clientUsecase) Update(id uint, name string) error {
	//Conver the data to model.Client struct
	client := &model.Client{
		Orm: model.Orm{
			ID: id,
		},
		Name: name,
	}
	return c.repo.Update(client)
}
