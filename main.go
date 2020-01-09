package main

import (
	"net/http"

	"github.com/dannywolfmx/ReSender/models"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	clientes := []models.Cliente{}

	server := gin.Default()

	//Archivos estaticos
	server.Use(static.Serve("/", static.LocalFile("./assets", true)))
	server.GET("/clientes", func(c *gin.Context) {
		c.JSON(http.StatusOK, clientes)
	})

	server.POST("/clientes", func(c *gin.Context) {
		var cliente models.Cliente
		if err := c.ShouldBindJSON(&cliente); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		clientes = append(clientes, cliente)

	})

	server.Run()
}
