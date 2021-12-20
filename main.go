package main

import (
	"url-shortener-api/config"
	"url-shortener-api/controller"
	"url-shortener-api/repository"
	"url-shortener-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database := config.NewSqliteDatabase()
	urlRepository := repository.NewUrlRepository(database)
	urlService := service.NewUrlService(urlRepository)
	urlController := controller.NewUrlController(urlService)
	// url := entity.Url{
	// 	Id:      shortuuid.New(),
	// 	LongUrl: "https://google.com",
	// 	UserId:  shortuuid.New(),
	// }
	// fmt.Println(urlService.Create(url))
	// // // err := urlRepository.Insert(url)
	// // // exception.PanicIfNeeded(err)
	// // // fmt.Println(url.Id)

	r := gin.Default()
	urlController.Route(r)
	r.Run(":8080")
}
