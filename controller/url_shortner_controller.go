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
	Fetch(ctx *gin.Context)
}

type urlShortnerController struct {
	service service.UrlShortnerService
}

func (u urlShortnerController) Fetch(ctx *gin.Context) {
	shortenedUrl := ctx.Param("shortenedUrl")
	if len(shortenedUrl) != 6 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println("fetching actual url")
	url := fmt.Sprint("http://localhost:8080/"+shortenedUrl)
	actualUrl, err := u.service.Fetch(url)
	if err != nil {
		fmt.Println("Error occurred while fetching actual url :: ", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, actualUrl)
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
