package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

type DataBase interface {
	InitDB() error
	Ping(data *sqlx.DB) error
}

type DBSqlite struct {
}

func (db DBSqlite) InitDB() (*sqlx.DB, error) {
	database, err := sqlx.Open("sqlite3", "./db/data/data.db")
	if err != nil {
		return nil, fmt.Errorf("Error al abrir la base de datos %g", err)
	}
	err = createTables(database)

	if err != nil {
		return nil, fmt.Errorf("Error al crear las tablas %g", err)
	}

	return database, nil
}

//Crear tablas por default en caso de que no existan
func createTables(db *sqlx.DB) error {
	//Crear Clientes
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS clientes (id TEXT PRIMARY KEY, nombre TEXT)")
	if err != nil {
		return fmt.Errorf("Error al crear query de tablas")
	}
	_, err = statement.Exec()
	if err != nil {
		return fmt.Errorf("Error ejecutar query de tablas")
	}

	return err
}

func (db DBSqlite) Ping(data *sqlx.DB) error {
	if err := data.Ping(); err != nil {
		fmt.Println("Error en Pong")
	} else {
		fmt.Println("Pong")
	}
	return nil
}
