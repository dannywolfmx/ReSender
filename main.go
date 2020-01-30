package main

import (
	"fmt"
	"log"

	"github.com/dannywolfmx/ReSender/db"
	"github.com/dannywolfmx/ReSender/models"
	"github.com/dannywolfmx/ReSender/repository"
	"github.com/rs/xid"
)

func main() {
	//Inicializar la base de datos
	sqlite := db.DBSqlite{}
	var err error
	db.DB, err = sqlite.InitDB()

	if err != nil {
		log.Fatalf("Error al momento de iniciar DB %g", err)
	}

	repositorio := repository.ClienteRepository{}

	cliente := &models.Cliente{Id: xid.New(), Nombre: "prueba2"}
	repositorio.Save(cliente)

	clientes := []models.Cliente{}
	repositorio.All(&clientes)

	fmt.Println(clientes)

	//	server := gin.Default()
	//	{
	//		route.Run(server)
	//	}
	//	server.Run()
}
