package repository

import (
	"fmt"

	"github.com/dannywolfmx/ReSender/db"
	"github.com/dannywolfmx/ReSender/models"
	"github.com/rs/xid"
)

//ClienteRepository representa un repositorio del tipo cliente
type ClienteRepository struct {
	Cliente *models.Cliente
}

//Save guarda un cliente
func (c *ClienteRepository) Save() error {
	query, err := db.DB.Prepare("INSERT INTO clientes (id, nombre) VALUES (?, ?)")

	if err != nil {
		return fmt.Errorf("Error al crear query de insertar cliente %g", err)
	}

	id := xid.New()
	_, err = query.Exec(id, c.Cliente.Nombre)

	if err != nil {
		return fmt.Errorf("Error al insertar un nuevo cliente %g", err)
	}

	return nil
}

//Get Retorna un cliente localizado por el id
func (c *ClienteRepository) Get(nombre string) (*models.Cliente, error) {
	filas, err := db.DB.Query("SELECT id, nombre FROM clientes")
	if err != nil {
		return nil, fmt.Errorf("Error al recuparar clientes %g", err)
	}

	for filas.Next() {
		filas.Scan(&id, &nombre)
		fmt.Println(nombre)
	}
	return nil, nil
}

//Delete elimina un cliente en la base de datos
//retorna el cliente encontrado y un estado de error en caso de existir este error
func (c *ClienteRepository) Delete(id xid.ID) (*models.Cliente, error) {
	cliente := &models.Cliente{}
	return cliente, nil
}

//All, ¡¡Importante, recoradar hacer una copia de lo que retorne, por la forma en que funciona Go puede mejorar el rendimiento!!
func (c *ClienteRepository) All() ([]models.Cliente, error) {
	filas, err := db.DB.Query("SELECT id, nombre FROM clientes")
	if err != nil {
		return nil, fmt.Errorf("Error al recuparar clientes %g", err)
	}
	var id xid.ID
	var nombre string
	for filas.Next() {
		filas.Scan(&id, &nombre)
		fmt.Println(nombre)
	}
	return nil, nil
}
