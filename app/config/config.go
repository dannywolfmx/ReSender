package config

import (
	"github.com/dannywolfmx/ReSender/app/delivery/http"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gin-gonic/gin"
)

func Init(routers *gin.RouterGroup) {
	container, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	http.Apply(routers, container)
}
