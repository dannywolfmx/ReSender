package http

import (
	v1 "github.com/dannywolfmx/ReSender/app/delivery/http/v1"
	"github.com/dannywolfmx/ReSender/registry"
	"github.com/gin-gonic/gin"
)

func Apply(router *gin.Engine, handler gin.HandlerFunc, ctn *registry.DIContainer) {
	api := router.Group("/api", handler)
	{
		v1.Apply(api, ctn)
	}
}
