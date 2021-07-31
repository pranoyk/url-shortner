package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlShortnerController interface {
	Shorten(ctx *gin.Context)
}

type urlShortnerController struct {}

func (u urlShortnerController) Shorten(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func NewUrlShortnerController() UrlShortnerController {
	return urlShortnerController{}
}