package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	"main/service"
	"main/util"
)

func Init() *gin.Engine {
	router := gin.Default()
	clock := util.NewClock()
	randomNumberGenerator := util.NewRandomNumberGenerator()
	randomStringGenerator := util.NewRandomStringGenerator(clock, randomNumberGenerator)
	shortnerService := service.NewUrlShortnerService(randomStringGenerator)
	urlShortnerController := controller.NewUrlShortnerController(shortnerService)
	router.POST("/api/url-shortner/v1/shorten", urlShortnerController.Shorten)
	return router
}
