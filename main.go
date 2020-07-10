package main

import (
	api "github.com/dannywolfmx/ReSender/app/delivery/http"
	appContainer "github.com/dannywolfmx/ReSender/app/registry"

	auth "github.com/dannywolfmx/ReSender/auth/delivery"
	authContainer "github.com/dannywolfmx/ReSender/auth/registry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	appCont, err := appContainer.NewContainer()
	if err != nil {
		panic(err)
	}

	authCont, err := authContainer.NewContainer()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	//Add cors for testing with localhost
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	api.Apply(router, appCont)
	auth.Apply(router, authCont)

	//Run the server
	router.Run(":8080")
}
