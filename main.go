package main

import (
	"github.com/dannywolfmx/ReSender/app/interface/api"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctn, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	//Add cors for testing with localhost
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	api.Apply(router, ctn)

	//Run the server
	router.Run(":8080")
}
