package application

import (
	"github.com/dannywolfmx/ReSender/app/interface/application/routes"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-gonic/gin"
)

func Apply(server *gin.Engine, ctn *registry.Container) {
	routes.Apply(server, ctn)
}
