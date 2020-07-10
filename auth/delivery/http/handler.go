package http

import (
	"fmt"
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

	err := h.u.SignUp(input.Username, input.Password)

	if err != nil {
		ctn.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctn.Status(http.StatusOK)
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
		}
		fmt.Println(err)
		ctn.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctn.JSON(http.StatusOK, signInResponseFields{Token: token})

}
