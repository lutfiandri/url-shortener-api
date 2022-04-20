package controller

import (
	"net/http"

	"url-shortener-api/config"
	"url-shortener-api/entity"
	"url-shortener-api/model"
	"url-shortener-api/service"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
)

type urlController struct {
	UrlService service.UrlService
	env        config.Env
}

func NewUrlController(urlService service.UrlService) UrlController {
	return &urlController{
		UrlService: urlService,
		env:        config.GetEnv(),
	}
}

func (c *urlController) Route(router *gin.Engine) {
	router.GET("/:id", c.FindByIdAndRedirect)
	router.GET("/urls/", c.FindAll)
	router.GET("/urls/:id", c.FindById)
	router.GET("/urls/by/:userId", c.FindByUserId)
	router.POST("/urls/", c.Create)
	router.DELETE("/urls/:id", c.DeleteById)
}

func (c *urlController) FindByIdAndRedirect(ctx *gin.Context) {
	id := ctx.Param("id")

	url, err := c.UrlService.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.Redirect(http.StatusMovedPermanently, url.LongUrl)
}

func (c *urlController) FindAll(ctx *gin.Context) {
	urls, err := c.UrlService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, urls)
}

func (c *urlController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	url, err := c.UrlService.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, model.UrlResponse{
		Id:       url.Id,
		ShortUrl: c.env.BaseUrl + "/" + url.Id,
		LongUrl:  url.LongUrl,
		UserId:   url.UserId,
	})
}

func (c *urlController) FindByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")

	urls, err := c.UrlService.FindByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, urls)
}

func (c *urlController) Create(ctx *gin.Context) {
	var request model.UrlRequestCreate
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	url := entity.Url{
		Id:      shortuuid.New(),
		LongUrl: request.LongUrl,
	}

	err = c.UrlService.Create(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, model.UrlResponse{
		Id:       url.Id,
		ShortUrl: c.env.BaseUrl + "/" + url.Id,
		LongUrl:  url.LongUrl,
		UserId:   url.UserId,
	})
}

func (c *urlController) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.UrlService.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
