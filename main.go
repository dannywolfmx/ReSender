package main

import (
	"fmt"
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

		if err := c.ShouldBind(&cliente); err != nil {
			c.String(http.StatusBadRequest, "Formato invalido")
			return
		}
		fmt.Println(cliente.Nombre)
		clientes = append(clientes, cliente)
		c.String(http.StatusOK, "ok")

	})

	server.Run()
}
