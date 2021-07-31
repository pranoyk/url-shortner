package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/service"
	"net/http"
)

type UrlShortnerController interface {
	Shorten(ctx *gin.Context)
}

type urlShortnerController struct {
	service service.UrlShortnerService
}

func (u urlShortnerController) Shorten(ctx *gin.Context) {
	fmt.Println("Shorten url initiated")
	shortenedUrl := u.service.Shorten()
	fmt.Println("url successfully shortened")
	ctx.JSON(http.StatusOK, shortenedUrl)
}

func NewUrlShortnerController(service service.UrlShortnerService) UrlShortnerController {
	return urlShortnerController{service: service}
}
