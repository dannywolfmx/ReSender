package api

import (
	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-gonic/gin"
)

func Apply(route *gin.Engine, ctn *registry.Container) {
	v1.Apply(route, ctn)
}
