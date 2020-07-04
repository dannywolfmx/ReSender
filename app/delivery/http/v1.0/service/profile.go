package service

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

type profileService struct {
	u usecase.ProfileUsecase
}

//NewProfileService devuelve un nuevo profileService con un usecase
func NewProfileService(u usecase.ProfileUsecase) *profileService {
	return &profileService{
		u: u,
	}
}

func (s *profileService) GetAll(c *gin.Context) {
	profiles, err := s.u.GetAll()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"erro": err,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		profiles,
	)
}

func (s *profileService) GetByID(c *gin.Context) {
	profileID, err := strconv.Atoi(c.Param("profileID"))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"erro": err,
			},
		)
		return
	}

	profile, err := s.u.GetByID(uint(profileID))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"erro": err,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		profile,
	)
}

//Create a new profile
func (s *profileService) Create(c *gin.Context) {
	//Create a profile data container
	profile := &usecase.Profile{}

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
				"error": err.Error(),
			},
		)
		return
	}
	c.JSON(
		http.StatusCreated,
		profile,
	)
}

//UpdatePassword
//func (s *profileService) UpdatePassword(c *gin.Context) {
//	profile := &usecase.ProfileWithPassowrd{}
//
//	//Bind the json information to the struct, and check if exist a error
//	if err := c.ShouldBind(profile); err != nil {
//		//Send a mmessege to the client with the error
//		c.JSON(
//			http.StatusBadRequest,
//			gin.H{
//				"code":  http.StatusBadRequest,
//				"error": "JSON invalid",
//			},
//		)
//		//Exit to the function
//		return
//	}
//
//	if err := s.u.UpdatePassword(profile.ProfileID, profile.Password); err != nil {
//		c.JSON(
//			http.StatusBadRequest,
//			gin.H{
//				"code":  http.StatusBadRequest,
//				"error": "error in the create funciton",
//			},
//		)
//		return
//	}
//	c.JSON(
//		http.StatusOK,
//		gin.H{
//			"code": http.StatusOK,
//		},
//	)
//}

//Delete a new profile
//TODO PROGRAMAR SISTEA DE AUTENTIFICACION
//func (s *profileService) Delete(c *gin.Context) {
//	//Create a profile data container
//	profile := &usecase.Profile{}
//
//	//Bind the json information to the struct, and check if exist a error
//	if err := c.ShouldBind(profile); err != nil {
//		//Send a mmessege to the client with the error
//		c.JSON(
//			http.StatusBadRequest,
//			gin.H{
//				"code":  http.StatusBadRequest,
//				"error": "JSON invalid",
//			},
//		)
//		//Exit to the function
//		return
//	}
//	if err := s.u.Delete(profile.ID); err != nil {
//		c.JSON(
//			http.StatusBadRequest,
//			gin.H{
//				"code":  http.StatusBadRequest,
//				"error": "error in the create funciton",
//			},
//		)
//		return
//	}
//	c.JSON(
//		http.StatusOK,
//		profile,
//	)
//}

//Update
func (s *profileService) Update(c *gin.Context) {
	//Create a profile data container
	profile := &usecase.Profile{}

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
	if err := s.u.Update(profile); err != nil {
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
		http.StatusOK,
		profile,
	)
}
