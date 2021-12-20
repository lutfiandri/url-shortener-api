package controller

import "github.com/gin-gonic/gin"

type UrlController interface {
	Route(r *gin.Engine)
	Create(c *gin.Context)
	GetLongUrl(c *gin.Context)
}