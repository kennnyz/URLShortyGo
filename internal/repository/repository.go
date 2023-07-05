package repository

import (
	"fmt"
	"ozonTech/muhtarov/internal/config"
	"ozonTech/muhtarov/internal/models"
	"ozonTech/muhtarov/internal/repository/inmemory_repository"
	"ozonTech/muhtarov/internal/repository/postgres_repository"
	"ozonTech/muhtarov/pkg/database/postgres"
)

//go:generate mockgen -source=repository.go -destination=mock/mock.go

type URLShortyRepository interface {
	AddUrl(urlStruct models.UrlStruct) (models.UrlStruct, error)
	GetFullUrlByShort(shortUrl string) (models.UrlStruct, error) // check if user exists
}

func NewRepository(cfg *config.Config) (URLShortyRepository, error) {
	switch cfg.StorageType {
	case "inmemory":
		return inmemory_repository.NewUrlShortRepo(), nil
	case "postgres":
		db, err := postgres.NewClient(cfg.DB.Dsn)
		if err != nil {
			panic(err)
		}
		return postgres_repository.NewEmailRepo(db), nil
	default:
		return nil, fmt.Errorf("unknown repository type")
	}
}
