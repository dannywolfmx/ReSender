package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	InitDB() error
}

type DBSqlite struct{}

func (db DBSqlite) InitDB() error {
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	err := createTables(database)
	if err != nil {
		return fmt.Errorf("Error al crear las tablas %g", err)
	}
	return nil
}

//Crear tablas por default en caso de que no existan
func createTables(db *sql.DB) error {
	//Crear Clientes
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS clientes (id TEXT PRIMARY KEY, nombre TEXT)")
	statement.Exec()
	return err
}
