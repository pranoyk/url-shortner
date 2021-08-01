package service

import (
	"errors"
	"fmt"
	"main/model"
	"main/util"
)

//go:generate mockgen -source=url_shortner_service.go -destination=../mocks/mock_url_shortner_service.go -package=mocks

type UrlShortnerService interface {
	Shorten(url string) model.ShortenResponseModel
	Fetch(shortenedUrl string) (string, error)
}

type urlShortnerService struct {
	randomStringGenerator util.RandomStringGenerator
	cacheUtil             util.CacheUtil
}

func (u urlShortnerService) Fetch(shortenedUrl string) (string, error){
	fmt.Println("performing fetch for url ", shortenedUrl)
	actualUrl, found := u.cacheUtil.Get(shortenedUrl)
	if found {
		fmt.Println("actual url found")
		return actualUrl.(string), nil
	}
	return "", errors.New("required url not found in our directory")
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
	u.cacheUtil.Set(newShortenedUrl, url)
	return model.ShortenResponseModel{
		ShortenedUrl: newShortenedUrl,
	}
}

func NewUrlShortnerService(randomStringGenerator util.RandomStringGenerator, cacheUtil util.CacheUtil) UrlShortnerService {
	return urlShortnerService{randomStringGenerator: randomStringGenerator, cacheUtil: cacheUtil}
}
