package service

import (
	"net/http"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

type profileService struct {
	u usecase.ProfileUsecase
}

func NewProficeService(u usecase.ProfileUsecase) *profileService {
	return &profileService{
		u: u,
	}
}

func (s *profileService) Create(c *gin.Context) {
	//Create a profile data container
	profile := &model.Profile{}

	//Bind the json information to the struct, and check if exist a error
	if err := c.ShouldBind(profile); err != nil {
		//Send a mmessege to the client with the error
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "JSON invalid",
			},
		)
		//Exit to the function
		return
	}
	if err := s.u.Create(profile); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "error in the create funciton",
			},
		)
		return
	}
	c.JSON(
		http.StatusCreated,
		profile,
	)
}
