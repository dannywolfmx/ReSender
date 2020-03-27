package main

import (
	"github.com/dannywolfmx/ReSender/app/interface/api"
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
	api.Apply(server, ctn)
	server.Run()
}
