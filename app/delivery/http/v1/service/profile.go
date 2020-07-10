package service

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app"
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

//Create a new profile
func (s *profileService) Create(c *gin.Context) {
	//Create a profile data container
	profile := &createProfile{}

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
	//Pass the struct data in "raw"
	if err := s.u.Create(profile.ImageAvatarPath, profile.Name, profile.Password); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			},
		)
		return
	}
	//Return the profile in JSON format
	c.JSON(
		http.StatusCreated,
		profile,
	)
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
	if err := s.u.Update(profile.ID, profile.ImageAvatarPath, profile.Name); err != nil {
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

//UpdatePassword
func (s *profileService) UpdatePassword(c *gin.Context) {
	profile := &updatePasswordProfile{}

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

	if err := s.u.UpdatePassword(profile.ProfileID, profile.Password); err != nil {
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
		gin.H{
			"code": http.StatusOK,
		},
	)
}
