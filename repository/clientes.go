package repository

import (
	"github.com/dannywolfmx/ReSender/models"
	"github.com/rs/xid"
)

//Repository cuenta con las funciones principales para manejar interactuar con un modelo en la base de datos
type Repository interface {
	Save() (bool, error)
	Get(id xid.ID) (*models.Cliente, error)
	Delete(id xid.ID) (*models.Cliente, error)
	All() ([]models.Cliente, error)
}

//ClienteRepository representa un repositorio del tipo cliente
type ClienteRepository struct{}

//Save guarda un cliente
func (c *ClienteRepository) save() (bool, error) {
	return true, nil
}

//Get Retorna un cliente localizado por el id
func (c *ClienteRepository) Get() (*models.Cliente, error) {
	cliente := &models.Cliente{}
	return cliente, nil
}

//Delete elimina un cliente en la base de datos
//retorna el cliente encontrado y un estado de error en caso de existir este error
func Delete(id xid.ID) (*models.Cliente, error) {
	cliente := &models.Cliente{}
	return cliente, nil
}

//All, ¡¡Importante, recoradar hacer una copia de lo que retorne, por la forma en que funciona Go puede mejorar el rendimiento!!
func All() ([]models.Cliente, error) {
	var clientes []models.Cliente
	return clientes, nil
}
