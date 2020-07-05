package service

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app"
	"github.com/gin-gonic/gin"
)

type orderService struct {
	u app.OrderUsecase
}

//NewOrderService construlle un servicio con un usecase
func NewOrderService(u app.OrderUsecase) *orderService {
	return &orderService{
		u: u,
	}
}

//Delete a element
func (s *orderService) Delete(c *gin.Context) {
	//Conver the id to numeric id
	id, err := strconv.Atoi(c.Param("id"))
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

	//Conver int to uint
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

type updateOrder struct {
	ID      uint   `json:"id"`
	Number  string `json:"number"`
	Invoice string `json:"invoice"`
}

func (s *orderService) Update(c *gin.Context) {
	order := &updateOrder{}
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

	err := s.u.Update(order.ID, order.Number, order.Invoice)
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
