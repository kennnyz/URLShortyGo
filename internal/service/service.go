package service

import (
	"ozonTech/muhtarov/internal/repository"
)

const base62Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type Service struct {
	UrlShortyService URLShorty
}

func NewService(repo repository.URLShortyRepository) *Service {
	return &Service{
		UrlShortyService: NewURLShortyService(repo),
	}
}
