package main

import (
	"url-shortener-api/config"
	"url-shortener-api/controller"
	"url-shortener-api/repository"
	"url-shortener-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database := config.NewSqliteDatabase(".data/url-shortener.db")
	urlRepository := repository.NewUrlRepository(database)
	urlService := service.NewUrlService(urlRepository)
	urlController := controller.NewUrlController(urlService)

	r := gin.Default()
	urlController.Route(r)
	r.Run(":8080")
}
