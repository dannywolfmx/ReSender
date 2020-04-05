package v1

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

func Apply(server *gin.Engine, ctn *registry.Container) {

	restClient(server, ctn)
	restOrder(server, ctn)
}

func restOrder(s *gin.Engine, ctn *registry.Container) {
	//Crear el caso de uso
	orderUseCase := NewOrderService(ctn.Resolve("order-usecase").(usecase.OrderUseCase))

	//List of orders
	s.GET("/orders", func(c *gin.Context) {
		j, err := orderUseCase.ListOrder()
		if err != nil {
			panic(err)
		}
		c.JSON(200, j)
	})

	s.POST("/order", func(c *gin.Context) {
		order := &model.Order{}
		c.BindJSON(order)
		err := orderUseCase.RegisterOrder(order.Number, order.Invoice, order.ClientID)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"order": order,
		})
	})
}

func restClient(s *gin.Engine, ctn *registry.Container) {
	clienteUseCase := NewClientService(ctn.Resolve("client-usecase").(usecase.ClientUseCase))
	//List of clients
	s.GET("/clients", func(c *gin.Context) {
		j, err := clienteUseCase.ListClient()
		if err != nil {
			panic(err)
		}
		c.JSON(200, j)
	})

	//Crear cliente
	s.POST("/client", func(c *gin.Context) {
		client := &model.Client{}
		c.BindJSON(client)
		err := clienteUseCase.RegisterClient(client.Name, client.Orders)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"client": client,
		})
	})
}