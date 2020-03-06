package v1

import (
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

func Apply(server *gin.Engine, ctn *registry.Container) {
	orderUseCase := NewOrderService(ctn.Resolve("order-usercase").(usecase.OrderUseCase))
	server.GET("/order", func(c *gin.Context) {
		j, err := orderUseCase.ListOrder()
		if err != nil {
			panic(err)
		}
		c.JSON(200, j)
	})
}
