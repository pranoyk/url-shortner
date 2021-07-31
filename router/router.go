package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func Init() *gin.Engine {
	router := gin.Default()
	urlShortnerController := controller.NewUrlShortnerController()
	router.POST("/api/url-shortner/v1/shorten", urlShortnerController.Shorten)
	return router
}

