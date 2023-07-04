package service

import (
	"hash/fnv"
	"math/big"
	"ozonTech/muhtarov/internal/models"
	"ozonTech/muhtarov/internal/repository"
)

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
		// TODO logg
		return models.UrlStruct{}, err
	}

	return m, nil
}

func (s *URLShortyService) makeShortURL(longURL string) models.UrlStruct {
	id := generateID(longURL)
	shortURL := encodeBase62(id)
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

func encodeBase62(number int64) string {

	base := big.NewInt(int64(len(base62Charset)))
	result := ""

	zero := big.NewInt(0)
	n := big.NewInt(number)
	for n.Cmp(zero) > 0 {
		quotient := new(big.Int)
		remainder := new(big.Int)
		quotient.DivMod(n, base, remainder)
		index := remainder.Int64()
		result = string(base62Charset[index]) + result
		n.Set(quotient)
	}

	return result
}
