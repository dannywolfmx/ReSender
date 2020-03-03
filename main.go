package main

import (
	"log"

	"github.com/dannywolfmx/ReSender/db"
)

//path sqlite

func main() {
	//Inicializar la base de datos
	pathDB := "./db/data/data.db"
	var err error
	db.DB, err = db.NewDBSqliteConnection(pathDB).InitDB()

	if err != nil {
		log.Fatalf("Error al momento de iniciar DB %g", err)
	}

	//	server := gin.Default()
	//	{
	//		route.Run(server)
	//	}
	//	server.Run()
}
