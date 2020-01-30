package repository

import (
	"github.com/dannywolfmx/ReSender/models"
	"github.com/rs/xid"
)

//Repository cuenta con las funciones principales para manejar interactuar con un modelo en la base de datos
type Repository interface {
	Save(cliente *models.Cliente) error
	Get(nombre string, cliente *models.Cliente) error
	Delete(id xid.ID) error
	All(clientes *[]models.Cliente) error
}
