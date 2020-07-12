//Entry point api
//Vercion 1 de la implementacion de la api Json

package v1

import (
	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/app/delivery/http/v1/service"
	"github.com/dannywolfmx/ReSender/registry"
	"github.com/gin-gonic/gin"
)

func Apply(r *gin.RouterGroup, ctn *registry.Container) {
	//Generate service with the use case
	client := service.NewClientService(ctn.Resolve("client-usecase").(app.ClientUsecase))
	order := service.NewOrderService(ctn.Resolve("order-usecase").(app.OrderUsecase))
	profile := service.NewProfileService(ctn.Resolve("profile-usecase").(app.ProfileUsecase))

	//REST SECTION

	//REST client
	r.POST("/client", client.Create)
	r.GET("/clients", client.List)
	r.DELETE("/client/:id", client.Delete)
	r.PUT("/client", client.Update)

	//REST orders
	r.PUT("/order", order.Update)
	r.DELETE("/order/:id", order.Delete)

	//REST Profile
	r.GET("/profiles", profile.GetAll)
	r.GET("/profile/:profileID", profile.GetAll)
	r.POST("/profile", profile.Create)
	r.PUT("/profile", profile.Update)
	r.PUT("/profile/updatePassword", profile.UpdatePassword)
	r.DELETE("/profile", profile.Delete)

}
