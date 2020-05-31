package route

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0/service"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

func Client(r *gin.Engine, ctn *registry.Container) {

	u := ctn.Resolve("client-usecase").(usecase.ClientUsecase)
	clientServiceApp := service.NewClientService(u)

	list := func(c *gin.Context) {
		clients, err := clientServiceApp.ListClient()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"code":  http.StatusInternalServerError,
					"error": "Error al buscar lista",
				},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			clients,
		)
	}

	create := func(c *gin.Context) {

		client := &model.Client{}
		if err := c.ShouldBind(client); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"code":  http.StatusBadRequest,
					"error": "JSON invalido",
				},
			)
			return

		}

		if err := clientServiceApp.RegisterClient(client); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"code":  http.StatusBadRequest,
					"error": "Error al crear",
				},
			)
			return
		}
		//Enviar respuesta de actualizacion exitoza
		c.JSON(
			http.StatusCreated,
			client,
		)
	}

	update := func(c *gin.Context) {

		client := &model.Client{}
		if err := c.ShouldBind(client); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"code":  http.StatusBadRequest,
					"error": "JSON invalido",
				},
			)
			return

		}

		if err := clientServiceApp.UpdateClient(client); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"code":  http.StatusBadRequest,
					"error": "Error al actualizar",
				},
			)
			return
		}
		//Enviar respuesta de actualizacion exitoza
		c.JSON(
			http.StatusCreated,
			client,
		)
	}

	//Delete a element
	remove := func(c *gin.Context) {
		idRemove := c.Param("id")

		//ID no numerico
		id, err := strconv.Atoi(idRemove)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"code":  http.StatusBadRequest,
					"error": "Id no numerico",
				},
			)
			return
		}

		//Cliente no encontrado
		if err := clientServiceApp.DeleteClient(uint(id)); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"code":  http.StatusBadRequest,
					"error": "Id no encontrado",
				},
			)
			return
		}
		//Enviar un mensaje de que se elimino de forma correcta
		c.JSON(
			http.StatusAccepted,
			gin.H{
				"code": http.StatusAccepted,
			},
		)
	}

	//Routes
	r.POST("/client", create)
	r.GET("/clients", list)
	r.DELETE("/client/:id", remove)
	r.PUT("/client", update)
}
