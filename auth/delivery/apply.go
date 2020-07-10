package delivery

import (
	"github.com/dannywolfmx/ReSender/auth"
	"github.com/dannywolfmx/ReSender/auth/delyvery/http"
	"github.com/dannywolfmx/ReSender/auth/registry"
	"github.com/gin-gonic/gin"
)

//Apply create a new endpoint
func Apply(router *gin.Engine, ctn *registry.Container) {
	http.RegisterHTTPEndPoint(router, ctn.Resolve("auth-usecase").(auth.AuthUsecase))
}
