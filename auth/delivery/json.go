package delivery

import (
	"github.com/dannywolfmx/ReSender/auth"
	"github.com/dannywolfmx/ReSender/auth/delivery/http"
	"github.com/dannywolfmx/ReSender/registry"
	"github.com/gin-gonic/gin"
)

func Apply(router *gin.Engine, ctn *registry.Container) {
	usecase := ctn.Resolve("auth-usecase").(auth.AuthUsecase)
	http.RegisterHTTPEndPoint(router, usecase)
}

func ApplyWithHandlers(router *gin.Engine, ctn *registry.Container, handler ...gin.HandlerFunc) {
	usecase := ctn.Resolve("auth-usecase").(auth.AuthUsecase)
	http.RegisterHTTPEndPoint(router, usecase)
}
