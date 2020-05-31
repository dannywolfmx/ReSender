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

//ORDERS REST
func Order(r *gin.Engine, ctn *registry.Container) {

	//Get a usecase from the container
	containerUsecase := ctn.Resolve("order-usecase").(usecase.OrderUseCase)
	//Crear el caso de uso
	orderServiceApp := service.NewOrderService(containerUsecase)

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

		//Orden no encontrada
		err = orderServiceApp.DeleteOrder(uint(id))
		if err != nil {
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

	update := func(c *gin.Context) {
		var order model.Order
		if err := c.ShouldBind(&order); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"code":  http.StatusBadRequest,
					"error": "JSON invalido",
				},
			)
			return

		}

		err := orderServiceApp.UpdateOrder(order.ClientID, order.Number, order.Invoice)
		if err != nil {
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
			http.StatusOK,
			order,
		)
	}

	//Routes
	r.PUT("/order", update)
	r.DELETE("/order/:id", remove)
}
