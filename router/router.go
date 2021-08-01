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
	cacheUtil := util.NewCacheUtil(cacheClient)
	clock := util.NewClock()
	randomNumberGenerator := util.NewRandomNumberGenerator()
	randomStringGenerator := util.NewRandomStringGenerator(clock, randomNumberGenerator)
	shortnerService := service.NewUrlShortnerService(randomStringGenerator, cacheUtil)
	urlShortnerController := controller.NewUrlShortnerController(shortnerService)

	router.POST("/api/url-shortner/v1/shorten", urlShortnerController.Shorten)
	router.GET("/:shortenedUrl",urlShortnerController.Fetch)
	return router
}
