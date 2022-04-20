package controller

import "github.com/gin-gonic/gin"

type UrlController interface {
	Route(r *gin.Engine)
	FindByIdAndRedirect(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	FindByUserId(c *gin.Context)
	Create(c *gin.Context)
	DeleteById(c *gin.Context)
}
