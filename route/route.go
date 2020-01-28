package route

import (
	"net/http"

	"github.com/dannywolfmx/ReSender/models"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var clientes []models.Cliente

func Run(server *gin.Engine) {
	initDB()
	//Archivos estaticos
	server.Use(static.Serve("/", static.LocalFile("./assets", true)))
	server.GET("/clientes", func(c *gin.Context) {
		c.JSON(http.StatusOK, clientes)
	})

	server.POST("/clientes", func(c *gin.Context) {
		var cliente models.Cliente
		//Verificar si los datos del formulario son correctos
		if err := c.ShouldBind(&cliente); err != nil {
			c.String(http.StatusBadRequest, "Datos invalidos")
			return
		}
		cliente.Id = xid.New()
		//Agregar a la db
		clientes = append(clientes, cliente)
		Guardar(clientes)
		//Enviar status de ok
		c.String(http.StatusOK, "ok")

	})

	server.Run()
}
