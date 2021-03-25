package main

import (
	"github.com/dannywolfmx/ReSender/config"
	"github.com/dannywolfmx/ReSender/server"
)

func main() {

	//TODO Hola
	//Get configuration
	configServer := config.ServerConfig()

	//Create server and run
	server.NewApp(configServer).Run()
}
