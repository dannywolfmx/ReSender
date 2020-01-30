package main

import (
	"log"

	"github.com/dannywolfmx/ReSender/db"
	"github.com/dannywolfmx/ReSender/models"
	"github.com/dannywolfmx/ReSender/repository"
)

func main() {
	//Inicializar la base de datos
	sqlite := db.DBSqlite{}
	var err error
	db.DB, err = sqlite.InitDB()

	if err != nil {
		log.Fatalf("Error al momento de iniciar DB %g", err)
	}

	repositorio := repository.ClienteRepository{
		Cliente: &models.Cliente{Nombre: "Prueba"},
	}
	repositorio.Save()

	repositorio.All()

	//	server := gin.Default()
	//	{
	//		route.Run(server)
	//	}
	//	server.Run()
}
