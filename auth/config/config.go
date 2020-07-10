package config

import (
	"github.com/dannywolfmx/ReSender/auth"
	"github.com/dannywolfmx/ReSender/auth/delivery/http"
	"github.com/dannywolfmx/ReSender/auth/registry"
	"github.com/gin-gonic/gin"
)

func Init(routers *gin.Engine) {
	authCont, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	authUsecase := authCont.Resolve("usecase").(auth.AuthUsecase)

	http.RegisterHTTPEndPoint(routers, authUsecase)
}
