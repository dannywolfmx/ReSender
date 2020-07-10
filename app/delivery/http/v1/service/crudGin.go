package service

import "github.com/gin-gonic/gin"

//CRUDGin es una interface para tomar la estrucutra de un crud realizado en Gin
type CRUDGin interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	All(c *gin.Context)
}
