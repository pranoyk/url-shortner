package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/model"
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
	var request model.ShortenUrlRequestModel
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println("Error occurred while binding JSON : ", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println("Shorten url initiated")
	shortenedUrl := u.service.Shorten(request.Url)
	fmt.Println("url successfully shortened")
	ctx.JSON(http.StatusOK, shortenedUrl)
}

func NewUrlShortnerController(service service.UrlShortnerService) UrlShortnerController {
	return urlShortnerController{service: service}
}
