package http

import (
	"github.com/dannywolfmx/ReSender/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndPoint(router *gin.Engine, authHandler auth.AuthUsecase) {
	handler := NewHandler(authHandler)
	group := router.Group("/auth")
	{
		group.POST("/sign-up", handler.SignUp)
		group.POST("/sign-in", handler.SignIn)
	}
}
