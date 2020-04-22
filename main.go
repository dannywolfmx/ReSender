package main

import (
	"github.com/dannywolfmx/ReSender/app/interface/application"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-gonic/gin"
)

//path sqlite

func main() {
	//Inicializar la base de datos

	ctn, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	server := gin.Default()
	application.Run(server, ctn)
	server.Run()
}
