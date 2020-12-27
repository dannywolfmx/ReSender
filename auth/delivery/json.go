package delivery

import (
	"github.com/dannywolfmx/ReSender/auth"
	"github.com/dannywolfmx/ReSender/auth/delivery/http"
	"github.com/gin-gonic/gin"
)

func Apply(router *gin.Engine, usecase auth.AuthUsecase) {
	http.RegisterHTTPEndPoint(router, usecase)
}

func ApplyWithHandlers(router *gin.Engine, usecase auth.AuthUsecase, handler ...gin.HandlerFunc) {
	http.RegisterHTTPEndPoint(router, usecase)
}
