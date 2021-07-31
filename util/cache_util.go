package util

import "github.com/patrickmn/go-cache"

//go:generate mockgen -source=cache_util.go -destination=../mocks/mock_cache_util.go -package=mocks

type CacheUtil interface {
	Set(key string, value string)
	Get(key string) (interface{}, bool)
}

type cacheUtil struct {
	cacheClient *cache.Cache
}

func (c cacheUtil) Get(key string) (interface{}, bool) {
	return c.cacheClient.Get(key)
}

func (c cacheUtil) Set(key string, value string) {
	c.cacheClient.Set(key, value, cache.NoExpiration)
}

func NewCacheUtil(cacheClient *cache.Cache) CacheUtil {
	return cacheUtil{cacheClient: cacheClient}
}
