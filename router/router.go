package router

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"main/controller"
	"main/service"
	"main/util"
)

func Init() *gin.Engine {
	router := gin.Default()

	cacheClient := cache.New(0, 0)
	clock := util.NewClock()
	randomNumberGenerator := util.NewRandomNumberGenerator()
	randomStringGenerator := util.NewRandomStringGenerator(clock, randomNumberGenerator)
	shortnerService := service.NewUrlShortnerService(randomStringGenerator, cacheClient)
	urlShortnerController := controller.NewUrlShortnerController(shortnerService)

	router.POST("/api/url-shortner/v1/shorten", urlShortnerController.Shorten)
	return router
}
