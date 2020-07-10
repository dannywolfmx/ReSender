package http

import (
	v1 "github.com/dannywolfmx/ReSender/app/delivery/http/v1"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-gonic/gin"
)

func Apply(router *gin.RouterGroup, ctn *registry.Container) {
	v1.Apply(router, ctn)
}
