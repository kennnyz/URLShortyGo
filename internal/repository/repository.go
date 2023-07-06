package repository

import (
	"fmt"
	"ozonTech/muhtarov/internal/config"
	"ozonTech/muhtarov/internal/repository/inmemory_repository"
	"ozonTech/muhtarov/internal/repository/postgres_repository"
	"ozonTech/muhtarov/internal/service"
	"ozonTech/muhtarov/pkg/database/postgres"
)

func NewRepository(cfg *config.Config) (service.URLShortyRepository, error) {
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
