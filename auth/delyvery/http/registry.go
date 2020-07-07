package http

import (
	"github.com/dannywolfmx/ReSender/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndPoint(router *gin.Engine, u auth.AuthUsecase) {
	handler := NewHandler(u)
	group := router.Group("/auth")
	{
		group.POST("/sign-up", handler.SignUp)
		group.POST("/sign-in", handler.SignIn)
	}
}
