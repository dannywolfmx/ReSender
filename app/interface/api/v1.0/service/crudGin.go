package service

import "github.com/gin-gonic/gin"

type CRUDGin interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	All(c *gin.Context)
}