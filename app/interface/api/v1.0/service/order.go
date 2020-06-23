package service

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

type orderService struct {
	u usecase.OrderUsecase
}

//NewOrderService construlle un servicio con un usecase
func NewOrderService(u usecase.OrderUsecase) *orderService {
	return &orderService{
		u: u,
	}
}

//Delete a element
func (s *orderService) Delete(c *gin.Context) {
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
	err = s.u.Delete(uint(id))
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

func (s *orderService) Update(c *gin.Context) {
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

	err := s.u.Update(order.ClientID, order.Number, order.Invoice)
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
