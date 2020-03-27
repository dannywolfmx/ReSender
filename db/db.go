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

type dbSqlite struct {
	path string
}

func NewDBSqliteConnection(path string) *dbSqlite {
	return &dbSqlite{
		path: path,
	}
}

func (db dbSqlite) Ping(data *sqlx.DB) error {
	if err := data.Ping(); err != nil {
		fmt.Println("Error en Pong")
	} else {
		fmt.Println("Pong")
	}
	return nil

}

func (db dbSqlite) InitDB() (*sqlx.DB, error) {
	database, err := sqlx.Open("sqlite3", db.path)
	if err != nil {
		return nil, err
	}
	err = createTables(database)

	if err != nil {
		return nil, err
	}

	return database, nil
}

var schemas = `
	CREATE TABLE IF NOT EXISTS orden (
		id TEXT PRIMARY KEY,
		number TEXT,
		invoice TEXT
	)
`

//createTables execute a schema creation if this exists
func createTables(db *sqlx.DB) error {
	//Create schemas
	_, err := db.Exec(schemas)
	if err != nil {
		return err
	}

	return nil
}
