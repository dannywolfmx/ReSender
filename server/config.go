package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	api "github.com/dannywolfmx/ReSender/app/delivery/http"
	"github.com/dannywolfmx/ReSender/registry"

	"github.com/dannywolfmx/ReSender/auth"
	authApi "github.com/dannywolfmx/ReSender/auth/delivery"
	"github.com/dannywolfmx/ReSender/auth/delivery/http"
)

type App struct {
	authUseCase auth.AuthUsecase
}

func NewApp() *App {
	authCont, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	return &App{
		authUseCase: authCont.Resolve("auth-usecase").(auth.AuthUsecase),
	}

}

func (a *App) Run(port string) {

	router := gin.Default()
	{
		a.initMiddleWare(router)
		a.initServices(router)
	}

	//Run the server
	router.Run(port)
}

func (a *App) initServices(router *gin.Engine) {
	authCont, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	authApi.Apply(router, authCont)

	container, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	api.Apply(router, http.NewAuthMiddleware(a.authUseCase), container)
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
