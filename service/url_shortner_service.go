package service

import (
	"fmt"
	"main/model"
	"main/util"
)

//go:generate mockgen -source=url_shortner_service.go -destination=../mocks/mock_url_shortner_service.go -package=mocks

type UrlShortnerService interface {
	Shorten(url string) model.ShortenResponseModel
}

type urlShortnerService struct {
	randomStringGenerator util.RandomStringGenerator
	cacheUtil             util.CacheUtil
}

func (u urlShortnerService) Shorten(url string) model.ShortenResponseModel {
	fmt.Println("performing shorten url in service")
	shortenedUrl, result := u.cacheUtil.Get(url)
	if result {
		fmt.Println("Found url in cache, returning from cache")
		return model.ShortenResponseModel{
			ShortenedUrl: shortenedUrl.(string),
		}
	}
	randString := u.randomStringGenerator.GetRandString(6)
	newShortenedUrl := fmt.Sprint("http://localhost:8080/" + randString)
	fmt.Println("Adding shortened url to cache")
	u.cacheUtil.Set(url, newShortenedUrl)
	return model.ShortenResponseModel{
		ShortenedUrl: newShortenedUrl,
	}
}

func NewUrlShortnerService(randomStringGenerator util.RandomStringGenerator, cacheUtil util.CacheUtil) UrlShortnerService {
	return urlShortnerService{randomStringGenerator: randomStringGenerator, cacheUtil: cacheUtil}
}
