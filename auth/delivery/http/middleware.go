package http

import (
	"net/http"
	"strings"

	"github.com/dannywolfmx/ReSender/auth"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	u auth.AuthUsecase
}

func NewAuthMiddleware(u auth.AuthUsecase) gin.HandlerFunc {
	midleware := &AuthMiddleware{
		u: u,
	}
	return midleware.Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	headers := strings.Split(authHeader, " ")
	if len(headers) != 2 || headers[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := headers[1]
	user, err := m.u.ParseToken(token)
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidToken {
			status = http.StatusUnauthorized
		}
		c.AbortWithStatus(status)
		return
	}

	c.Set(auth.ContexUserKey, user)

}
