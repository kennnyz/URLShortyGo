package repository

import (
	"database/sql"
	"fmt"
	"ozonTech/muhtarov/internal/models"
	"ozonTech/muhtarov/internal/repository/inmemory_repository"
	"ozonTech/muhtarov/internal/repository/postgres_repository"
)

type URLShortyRepository interface {
	AddUrl(urlStruct models.UrlStruct) (models.UrlStruct, error)
	GetFullUrlByShort(shortUrl string) (models.UrlStruct, error) // check if user exists
}

func NewRepository(repoType string, db *sql.DB) (URLShortyRepository, error) {
	switch repoType {
	case "inmemory":
		return inmemory_repository.NewUrlShortRepo(), nil
	case "postgres":
		if db == nil {
			return nil, fmt.Errorf("db is nil")
		}
		return postgres_repository.NewEmailRepo(db), nil
	default:
		return nil, fmt.Errorf("unknown repository type")
	}
}
