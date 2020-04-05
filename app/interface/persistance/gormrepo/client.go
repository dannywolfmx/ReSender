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

func (r clientRepository) FindByName(name string) (*model.Client, error) {
	client := new(model.Client)
	r.db.Where(&model.Client{Name: name}).First(client)
	return client, nil
}

//TODO: Convertir esta lista a una lista de apuntadores "[]*model.Client"
func (r clientRepository) All() ([]model.Client, error) {
	clients := []model.Client{}
	//Pedir a GORM que agregue las ordenes del usuario
	r.db.Preload("Orders").Find(&clients)
	return clients, nil
}

//Find and Delete all the matches record
//Note: Delete is a soft delete, this function just set a flag
//You need to use r.db.Unscoped().Delete(&model.Order{}) to clear the Delete records permanently
func (r clientRepository) Detele(name string) error {
	r.db.Where("name = ?", name).Delete(&model.Client{})
	return nil
}

func (r clientRepository) Update(client *model.Client) error {
	//Save will update all the fields, even it is not changed
	r.db.Save(client)
	return nil
}