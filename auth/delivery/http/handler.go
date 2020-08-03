package http

import (
	"net/http"

	"github.com/dannywolfmx/ReSender/auth"
	"github.com/gin-gonic/gin"
)

type handler struct {
	u auth.AuthUsecase
}

func NewHandler(u auth.AuthUsecase) *handler {
	return &handler{
		u: u,
	}
}

type signFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) SignUp(ctn *gin.Context) {
	input := &signFields{}
	if err := ctn.BindJSON(input); err != nil {
		ctn.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Create user and get the JWT token
	token, err := h.u.SignUp(input.Username, input.Password)

	if err != nil {
		if err == auth.ErrInvalidToken {
			ctn.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err == auth.ErrNameAlreayExist {
			ctn.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": auth.ErrNameAlreayExist.Error(),
			})
			return
		}
		ctn.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctn.JSON(http.StatusCreated, signInResponseFields{Token: token})
}

type signInResponseFields struct {
	Token string `json:"token"`
}

func (h *handler) SignIn(ctn *gin.Context) {
	input := &signFields{}

	if err := ctn.BindJSON(input); err != nil {
		ctn.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.u.SignIn(input.Username, input.Password)
	if err != nil {
		if err == auth.ErrInvalidToken {
			ctn.AbortWithStatus(http.StatusUnauthorized)
			return
		} else if err == auth.ErrInvalidPassword {
			ctn.AbortWithStatus(http.StatusUnauthorized)
			return
		} else if err == auth.ErrInvalidUser {
			ctn.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctn.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctn.JSON(http.StatusOK, signInResponseFields{Token: token})

}
