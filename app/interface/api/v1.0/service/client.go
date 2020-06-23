package service

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gin-gonic/gin"
)

type clientService struct {
	u usecase.ClientUsecase
}

//NewClientService genera un nuevo servicio de tipo client con un usecase
func NewClientService(u usecase.ClientUsecase) *clientService {
	return &clientService{
		u: u,
	}
}

func (s *clientService) List(c *gin.Context) {
	clients, err := s.u.Clients()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Error al buscar lista",
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		clients,
	)
}

func (s *clientService) Create(c *gin.Context) {

	client := &model.Client{}

	if err := c.ShouldBind(client); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "JSON invalid",
			},
		)
		return

	}

	if err := s.u.Register(client); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "Error al crear",
			},
		)
		return
	}
	//Enviar respuesta de actualizacion exitoza
	c.JSON(
		http.StatusCreated,
		client,
	)
}

func (s *clientService) Update(c *gin.Context) {

	client := &model.Client{}
	if err := c.ShouldBind(client); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "JSON invalido",
			},
		)
		return

	}

	if err := s.u.Update(client); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "Error al actualizar",
			},
		)
		return
	}
	//Enviar respuesta de actualizacion exitoza
	c.JSON(
		http.StatusCreated,
		client,
	)
}

//Delete a element
func (s *clientService) Delete(c *gin.Context) {
	idRemove := c.Param("id")

	//ID no numerico
	id, err := strconv.Atoi(idRemove)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "Id no numerico",
			},
		)
		return
	}

	//Cliente no encontrado
	if err := s.u.Delete(uint(id)); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":  http.StatusBadRequest,
				"error": "Id no encontrado",
			},
		)
		return
	}
	//Enviar un mensaje de que se elimino de forma correcta
	c.JSON(
		http.StatusAccepted,
		gin.H{
			"code": http.StatusAccepted,
		},
	)
}
