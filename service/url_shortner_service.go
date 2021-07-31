package service

import (
	"fmt"
	"main/model"
	"main/util"
)

type UrlShortnerService interface {
	Shorten() model.ShortenResponseModel
}

type urlShortnerService struct {
	randomStringGenerator util.RandomStringGenerator
}

func (u urlShortnerService) Shorten() model.ShortenResponseModel {
	fmt.Println("performing shorten url in service")
	randString := u.randomStringGenerator.GetRandString(6)
	return model.ShortenResponseModel{
		ShortenedUrl: randString,
	}
}

func NewUrlShortnerService(randomStringGenerator util.RandomStringGenerator) UrlShortnerService {
	return urlShortnerService{randomStringGenerator: randomStringGenerator}
}
