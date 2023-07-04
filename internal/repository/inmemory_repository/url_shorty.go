package inmemory_repository

import (
	"ozonTech/muhtarov/internal/models"
	"sync"
)

type UrlRepoInMemory struct {
	shortedToLong map[string]models.UrlStruct
	longToShorted map[string]models.UrlStruct
	mutex         *sync.RWMutex
}

func NewUrlShortRepo() *UrlRepoInMemory {
	return &UrlRepoInMemory{
		shortedToLong: make(map[string]models.UrlStruct),
		longToShorted: make(map[string]models.UrlStruct),
	}
}

func (e *UrlRepoInMemory) AddUrl(urlStruct models.UrlStruct) (models.UrlStruct, error) {

	if _, ok := e.longToShorted[urlStruct.LongUrl]; ok {
		return e.longToShorted[urlStruct.LongUrl], nil
	}

	e.mutex.Lock()
	e.longToShorted[urlStruct.LongUrl] = urlStruct
	e.shortedToLong[urlStruct.ShortUrl] = urlStruct
	e.mutex.Unlock()

	return urlStruct, nil
}

func (e *UrlRepoInMemory) GetFullUrlByShort(shortUrl string) (models.UrlStruct, error) {

	if _, ok := e.shortedToLong[shortUrl]; !ok {
		return models.UrlStruct{}, models.UrlNotFoundErr
	}

	return e.shortedToLong[shortUrl], nil
}
