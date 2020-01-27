package main

import (
	"github.com/dannywolfmx/ReSender/route"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	{
		route.Run(server, db)
	}
	server.Run()
}
