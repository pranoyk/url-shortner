package service

import (
	"fmt"
	"main/model"
	"main/util"
)

//go:generate mockgen -source=url_shortner_service.go -destination=../mocks/mock_url_shortner_service.go -package=mocks

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
		ShortenedUrl: "/"+randString,
	}
}

func NewUrlShortnerService(randomStringGenerator util.RandomStringGenerator) UrlShortnerService {
	return urlShortnerService{randomStringGenerator: randomStringGenerator}
}
