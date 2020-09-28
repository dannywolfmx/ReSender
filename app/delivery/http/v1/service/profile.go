package service

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app"
	"github.com/dannywolfmx/ReSender/auth/domain/model"
	"github.com/gin-gonic/gin"
)

type profileService struct {
	u app.ProfileUsecase
}

//NewProfileService devuelve un nuevo profileService con un usecase
func NewProfileService(u app.ProfileUsecase) *profileService {
	return &profileService{
		u: u,
	}
}

type UserContext struct {
	Username string
	Password string
}

//GetByContext get the profile data from the auth context
func (s *profileService) GetByContext(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	profile, err := s.u.GetByUserID(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"erro": "Error al obtener el usuario",
		})
		return
	}

	//The profile dont exist, we need to create a new one
	if profile == nil {
		//Set the new profile
		profile, err = s.u.Create(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"erro": "Error al tratar de crear un nuevo profile",
			})
			return
		}
	}

	c.JSON(http.StatusOK, profile)
}

func (s *profileService) GetAll(c *gin.Context) {
	//Get a []*model.Profile
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

	//Return the profiles in JSON format
	c.JSON(
		http.StatusOK,
		profiles,
	)
}

func (s *profileService) GetByID(c *gin.Context) {
	//Get and convert the "profileID" param to int
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

	//Convert the profileID int to uint
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

	//Return the profile in JSON format
	c.JSON(
		http.StatusOK,
		profile,
	)
}

//Use this struct to get the JSON data
type createProfile struct {
	ImageAvatarPath string `json:"image_avatar_path"`
	Name            string `json:"name"`
	Password        string `json:"password"`
}

type deleteProfile struct {
	ProfileID uint `json:"profile_id"`
}

//Delete a new profile
//TODO PROGRAMAR SISTEA DE AUTENTIFICACION
func (s *profileService) Delete(c *gin.Context) {
	//Create a profile data container
	profile := &deleteProfile{}

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
	if err := s.u.Delete(profile.ProfileID); err != nil {
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

type updateProfile struct {
	ID              uint   `json:"id"`
	ImageAvatarPath string `json:"image_avatar_path"`
	Name            string `json:"name"`
	Password        string `json:"password"`
}

//Update
func (s *profileService) Update(c *gin.Context) {
	//Create a profile data container
	profile := &updateProfile{}

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
	if err := s.u.Update(profile.ID, profile.ImageAvatarPath); err != nil {
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

type updatePasswordProfile struct {
	ProfileID uint   `json:"profile_id"`
	Password  string `json:"password"`
}
