package routes

import (
	"net/http"

	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

func Apply(server *gin.Engine, ctn *registry.Container) {
	server.LoadHTMLGlob("templates/**/*.tmpl")
	index(server, ctn)
}

func index(s *gin.Engine, ctn *registry.Container) {
	clienteUseCase := v1.NewClientService(ctn.Resolve("client-usecase").(usecase.ClientUseCase))
	s.GET("/", func(c *gin.Context) {
		j, err := clienteUseCase.ListClient()
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Clientes": j,
		})
	})
}
