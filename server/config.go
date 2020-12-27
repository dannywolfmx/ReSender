package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	api "github.com/dannywolfmx/ReSender/app/delivery/http"
	"github.com/dannywolfmx/ReSender/config"
	"github.com/dannywolfmx/ReSender/registry"

	authApi "github.com/dannywolfmx/ReSender/auth/delivery"
	"github.com/dannywolfmx/ReSender/auth/delivery/http"
)

type App struct {
	diContainer *registry.DIContainer
	config *config.Server
}

//NewApp create a new app with the app information
func NewApp(config *config.Server) *App {
	diContainer, err := registry.NewDIContainer(config.DbType)
	if err != nil {
		panic(err)
	}
	return &App{
		diContainer: diContainer,
		config:config,
	}

}

//Run the server
func (a *App) Run() {

	router := gin.Default()
	{
		a.initMiddleWare(router)
		a.initServices(router)
	}

	port := fmt.Sprintf(":%d", a.config.Port)

	//Run the server
	router.Run(port)
}

func (a *App) initServices(router *gin.Engine) {
	authUsecase := a.diContainer.AuthUsecase
	handler := http.NewAuthMiddleware(authUsecase)
	authApi.Apply(router,authUsecase)
	api.Apply(router,handler, a.diContainer)
}

func (a *App) initMiddleWare(router *gin.Engine) {
	setCORS(router, "*")
	//setCORS(router, "http://localhost:3000")
}

func setCORS(router *gin.Engine, origins ...string) {
	//Add cors for testing with localhost
	config := cors.DefaultConfig()
	config.AllowOrigins = origins

	router.Use(cors.New(config))
}
