package main

import (
	"github.com/dannywolfmx/ReSender/app/interface/api"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-gonic/gin"
)

func main() {
	ctn, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	route := gin.Default()
	api.Apply(route, ctn)

	//Run the server
	route.Run(":8080")
}
