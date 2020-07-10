package main

import (
	app "github.com/dannywolfmx/ReSender/app/config"
	"github.com/dannywolfmx/ReSender/auth"
	authConfig "github.com/dannywolfmx/ReSender/auth/config"
	"github.com/dannywolfmx/ReSender/auth/delivery/http"
	"github.com/dannywolfmx/ReSender/auth/registry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Configuration struct {
}

func main() {

	router := gin.Default()

	//Add cors for testing with localhost
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))
	authConfig.Init(router)

	authCont, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	authUsecase := authCont.Resolve("usecase").(auth.AuthUsecase)

	api := router.Group("/api", http.NewAuthMiddleware(authUsecase))
	app.Init(api)

	//Run the server
	router.Run(":8080")
}
