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
func (c *ClienteRepository) Save(cliente *models.Cliente) error {
	query, err := db.DB.Prepare("INSERT INTO clientes (id, nombre) VALUES (?, ?)")

	if err != nil {
		return fmt.Errorf("Error al crear query de insertar cliente %g", err)
	}

	_, err = query.Exec(cliente.Id, cliente.Nombre)

	if err != nil {
		return fmt.Errorf("Error al insertar un nuevo cliente %g", err)
	}

	return nil
}

//Get Retorna un cliente localizado por el id
func (c *ClienteRepository) Get(nombre string, cliente *models.Cliente) error {
	filas, err := db.DB.Query("SELECT id, nombre FROM clientes")
	if err != nil {
		return fmt.Errorf("Error al recuparar clientes %g", err)
	}

	var id xid.ID
	var name string
	for filas.Next() {
		filas.Scan(&id, &name)
		fmt.Println(name)
	}
	return nil
}

//Delete elimina un cliente en la base de datos
//retorna el cliente encontrado y un estado de error en caso de existir este error
func (c *ClienteRepository) Delete(id xid.ID) error {
	//cliente := &models.Cliente{}
	return nil
}

//All genera una lista de clientes
func (c *ClienteRepository) All(clientes *[]models.Cliente) error {
	filas, err := db.DB.Query("SELECT id, nombre FROM clientes")
	if err != nil {
		return fmt.Errorf("Error al recuparar clientes %g", err)
	}
	var id xid.ID
	var nombre string
	for filas.Next() {
		filas.Scan(&id, &nombre)
		cliente := models.Cliente{id, nombre}
		*clientes = append(*clientes, cliente)
	}
	return nil
}
