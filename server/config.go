package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dannywolfmx/ReSender/auth"
	"github.com/dannywolfmx/ReSender/auth/delivery/http"
	"github.com/dannywolfmx/ReSender/auth/registry"

	appRouter "github.com/dannywolfmx/ReSender/app/config"
	authRouter "github.com/dannywolfmx/ReSender/auth/config"
	authContainer "github.com/dannywolfmx/ReSender/auth/registry"
)

type App struct {
	authUseCase auth.AuthUsecase
}

func NewApp() *App {
	authCont, err := authContainer.NewContainer()
	if err != nil {
		panic(err)
	}
	defaultApp := &App{
		authUseCase: authCont.Resolve("usecase").(auth.AuthUsecase),
	}

	return defaultApp
}

func (a *App) Run() {

	router := gin.Default()
	{
		initMiddleWare(router)
		initServices(router)
	}

	//Run the server
	router.Run(":8080")
}

func initServices(router *gin.Engine) {
	authRouter.Init(router)

	authCont, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	authUsecase := authCont.Resolve("usecase").(auth.AuthUsecase)
	api := router.Group("/api", http.NewAuthMiddleware(authUsecase))
	{
		appRouter.Init(api)
	}
}

func initMiddleWare(router *gin.Engine) {
	//	setCORS(router, "http://localhost:3000")
}

func setCORS(router *gin.Engine, origins ...string) {
	//Add cors for testing with localhost
	config := cors.DefaultConfig()
	config.AllowOrigins = origins

	router.Use(cors.New(config))
}
