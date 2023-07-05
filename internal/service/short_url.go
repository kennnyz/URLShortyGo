package service

import (
	"hash/fnv"
	"ozonTech/muhtarov/internal/models"
	"ozonTech/muhtarov/internal/repository"
)

//go:generate mockgen -source=short_url.go -destination=mock/mock.go
type URLShorty interface {
	AddUrl(mail string) (models.UrlStruct, error)
	GetFullUrl(keyword string) (models.UrlStruct, error) // check if user exists
}

type URLShortyService struct {
	repo repository.URLShortyRepository
}

func NewURLShortyService(repo repository.URLShortyRepository) *URLShortyService {
	return &URLShortyService{
		repo: repo,
	}
}

func (s *URLShortyService) AddUrl(longUrl string) (models.UrlStruct, error) {
	shortUrl := s.makeShortURL(longUrl)
	model, err := s.repo.AddUrl(shortUrl)
	if err != nil {
		return models.UrlStruct{}, err
	}

	return model, nil
}

func (s *URLShortyService) GetFullUrl(keyword string) (models.UrlStruct, error) {
	m, err := s.repo.GetFullUrlByShort(keyword)
	if err != nil {
		return models.UrlStruct{}, err
	}

	return m, nil
}

func (s *URLShortyService) makeShortURL(longURL string) models.UrlStruct {
	id := generateID(longURL)
	shortURL := encodeBase62(id)
	if len(shortURL) < 10 {
		for i := 0; i <= 10-len(shortURL); i++ {
			shortURL = "0" + shortURL
		}
	}
	return models.UrlStruct{
		ShortUrl: shortURL,
		LongUrl:  longURL,
		Id:       id,
	}
}

// generateID generates a unique ID for a given string
func generateID(longURL string) int64 {
	h := fnv.New64a()
	h.Write([]byte(longURL))

	id := int64(h.Sum64())

	// making len of id 10 digits
	id = id % 100000000000000000
	if id < 0 {
		return -1 * id
	}
	return id
}

func encodeBase62(num int64) string {
	if num == 0 {
		return string(base62Charset[0])
	}

	if num < 0 {
		return "-" + encodeBase62(-num)
	}

	base62 := make([]byte, 0)
	for num > 0 {
		remainder := num % 62
		base62 = append([]byte{base62Charset[remainder]}, base62...)
		num /= 62
	}

	return string(base62)
}
