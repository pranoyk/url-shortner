package service

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"main/model"
	"main/util"
)

//go:generate mockgen -source=url_shortner_service.go -destination=../mocks/mock_url_shortner_service.go -package=mocks

type UrlShortnerService interface {
	Shorten(url string) model.ShortenResponseModel
}

type urlShortnerService struct {
	randomStringGenerator util.RandomStringGenerator
	cacheClient           *cache.Cache
}

func (u urlShortnerService) Shorten(url string) model.ShortenResponseModel {
	fmt.Println("performing shorten url in service")
	shortenedUrl, result := u.cacheClient.Get(url)
	if result {
		fmt.Println("Found url in cache, returning from cache")
		return model.ShortenResponseModel{
			ShortenedUrl: shortenedUrl.(string),
		}
	}
	randString := u.randomStringGenerator.GetRandString(6)
	newShortenedUrl := fmt.Sprint("localhost:8080/" + randString)
	fmt.Println("Adding shortened url to cache")
	u.cacheClient.Set(url, newShortenedUrl, cache.DefaultExpiration)
	return model.ShortenResponseModel{
		ShortenedUrl: newShortenedUrl,
	}
}

func NewUrlShortnerService(randomStringGenerator util.RandomStringGenerator, cacheClient *cache.Cache) UrlShortnerService {
	return urlShortnerService{randomStringGenerator: randomStringGenerator, cacheClient: cacheClient}
}
