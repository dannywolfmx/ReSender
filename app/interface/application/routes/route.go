package routes

import (
	"net/http"

	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-gonic/gin"
)

func Apply(server *gin.Engine, ctn *registry.Container) {
	server.LoadHTMLGlob("./templates/*")
	index(server, ctn)
}

func index(s *gin.Engine, ctn *registry.Container) {
	s.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":  "Pagina de inicio",
			"prueba": "Pagina de inicio",
		})
	})
}
