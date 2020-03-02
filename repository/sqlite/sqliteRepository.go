package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/dannywolfmx/ReSender/db"
	"github.com/dannywolfmx/ReSender/models"
	"github.com/dannywolfmx/ReSender/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type sqliteRepository struct {
	client   *sqlx.DB
	database string
	timeout  time.Duration
}

func newSqliteClient(sqliteURL string, sqliteTimeOut int) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(sqliteTimeOut)*time.Second)
	defer cancel()
	client, err := sqlx.Open("sqlite3", sqliteURL)
	if err != nil {
		return nil, err
	}

	err = client.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return client, err
}

//NewSqliteRepository Generate a new connection to the database Sqlite
func NewSqliteClient(sqliteURL, sqliteDB string, sqliteTimeOut int) (repository.Client, error) {
	repo := &sqliteRepository{
		timeout:  time.Duration(sqliteTimeOut) * time.Second,
		database: sqliteDB,
	}
	client, err := newSqliteClient(sqliteURL, sqliteTimeOut)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewSqliteRepository")
	}
	repo.client = client

	return repo, nil
}

//Save guarda un cliente
func (c *sqliteRepository) Save(client *models.Client) error {
	query, err := db.DB.Prepare("INSERT INTO clientes (id, nombre) VALUES (?, ?)")

	if err != nil {
		return fmt.Errorf("Error al crear query de insertar cliente %g", err)
	}

	_, err = query.Exec(client.ID, client.Name)

	if err != nil {
		return fmt.Errorf("Error al insertar un nuevo cliente %g", err)
	}

	return nil
}

//Get Retorna un cliente localizado por el id
func (c *sqliteRepository) Get(name string, client *models.Client) error {
	filas, err := db.DB.Query("SELECT id, nombre FROM clientes")
	if err != nil {
		return fmt.Errorf("Error al recuparar clientes %g", err)
	}

	var id xid.ID
	for filas.Next() {
		filas.Scan(&id, &name)
	}
	return nil
}

//Delete elimina un cliente en la base de datos
//retorna el cliente encontrado y un estado de error en caso de existir este error
func (c *sqliteRepository) Delete(id xid.ID) error {
	//cliente := &models.Cliente{}
	return nil
}

//All genera una lista de clientes
func (c *sqliteRepository) All(clients *[]models.Client) error {
	filas, err := db.DB.Query("SELECT id, nombre FROM clientes")
	if err != nil {
		return fmt.Errorf("Error al recuparar clientes %g", err)
	}
	var id xid.ID
	var name string
	for filas.Next() {
		filas.Scan(&id, &name)
		client := models.Client{ID: id, Name: name}
		*clients = append(*clients, client)
	}
	return nil
}
