package main

import (
	api "github.com/dannywolfmx/ReSender/app/delivery/http"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	appContainer, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	//Add cors for testing with localhost
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	api.Apply(router, appContainer)

	//Run the server
	router.Run(":8080")
}
