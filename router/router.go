package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()
	registerController(router)
	return router
}

func registerController(router *gin.Engine) {
	router.GET("/api/v1/url-shortner", shortner)
}
func shortner(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusOK)
}
