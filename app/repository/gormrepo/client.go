package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *clientRepository {
	return &clientRepository{
		db: db,
	}
}

func (r clientRepository) Save(client *model.Client) error {
	r.db.Save(client)
	return nil
}

func (r clientRepository) GetById(id uint) model.Client {
	client := model.Client{}
	r.db.Preload("Orders.Mails").Where("id = ?", id).First(&client)
	return client
}

//TODO: Convertir esta lista a una lista de apuntadores "[]*model.Client"
func (r *clientRepository) All() ([]*model.Client, error) {
	clients := []*model.Client{}
	//Pedir a GORM que agregue las ordenes del usuario
	//r.db.Set("gorm:auto_preload", true).Find(&clients)
	r.db.Preload("Orders.Files").Preload("Orders.Mails").Find(&clients)
	return clients, nil
}

//Find and Delete all the matches record
//Note: Delete is a soft delete, this function just set a flag
//You need to use r.db.Unscoped().Delete(&model.Order{}) to clear the Delete records permanently
func (r *clientRepository) Detele(id uint) error {
	r.db.Where("id = ?", id).Delete(&model.Client{})
	return nil
}

func (r clientRepository) Update(client *model.Client) error {
	//Save will update all the fields, even it is not changed
	r.db.Save(client)
	return nil
}

//Find a client by id
//If the client dont exists the struct is nil
func (r *clientRepository) Find(id uint) (*model.Client, error) {
	client := &model.Client{}
	r.db.Where("id = ?", id).First(client)
	return client, nil
}

//FindByName a client by name
//If the client dont exists the struct is nil
func (r *clientRepository) FindByName(name string) (*model.Client, error) {
	client := &model.Client{}
	r.db.Where("name = ?", name).First(client)
	return client, nil
}
