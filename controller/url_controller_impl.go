package controller

import (
	"net/http"
	"url-shortener-api/entity"
	"url-shortener-api/exception"
	"url-shortener-api/model"
	"url-shortener-api/service"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
)

type urlController struct {
	UrlService service.UrlService
}

func NewUrlController(urlService service.UrlService) UrlController {
	return &urlController{
		UrlService: urlService,
	}
}

func (c *urlController) Route(router *gin.Engine) {
	router.POST("/", c.Create)
	router.GET("/:id", c.GetLongUrl)
}

func (c *urlController) Create(ctx *gin.Context) {
	var request model.UrlRequestCreate
	err := ctx.BindJSON(&request)
	exception.PanicIfNeeded(err)

	url := entity.Url{
		Id:      shortuuid.New(),
		LongUrl: request.LongUrl,
	}

	c.UrlService.Create(url)

	shortUrl := "http://localhost:8080/" + url.Id

	ctx.JSON(http.StatusCreated, gin.H{
		"status":    "success",
		"short_url": shortUrl,
	})
}

func (c *urlController) GetLongUrl(ctx *gin.Context) {
	id := ctx.Param("id")

	longUrl := c.UrlService.GetLongUrl(id)

	ctx.Redirect(http.StatusMovedPermanently, longUrl)

	// ctx.JSON(http.StatusCreated, gin.H{
	// 	"status":   "success",
	// 	"long_url": longUrl,
	// })
}
