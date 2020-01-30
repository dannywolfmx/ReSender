package repository

import (
	"github.com/dannywolfmx/ReSender/models"
	"github.com/rs/xid"
)

//Repository cuenta con las funciones principales para manejar interactuar con un modelo en la base de datos
type Repository interface {
	Save() error
	Get(id xid.ID) (*models.Cliente, error)
	Delete(id xid.ID) (*models.Cliente, error)
	All() ([]models.Cliente, error)
}
