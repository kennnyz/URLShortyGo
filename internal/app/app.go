package app

import (
	"fmt"
	"log"
	"ozonTech/muhtarov/internal/config"
	"ozonTech/muhtarov/internal/delivery"
	"ozonTech/muhtarov/internal/repository"
	"ozonTech/muhtarov/internal/server"
	"ozonTech/muhtarov/internal/service"
	"ozonTech/muhtarov/pkg/database/postgres"
)

func Run(configPath string) {
	cfg, err := config.ReadConfig(configPath)
	if err != nil {
		panic(err)
	}

	db, err := postgres.NewClient(cfg.DB.Dsn)
	if err != nil {
		panic(err)
	}

	repos, err := repository.NewRepository(cfg.StorageType, db)
	if err != nil {
		panic(err)
	}

	services := service.NewService(repos)
	handler := delivery.NewHandler(services)

	httpServer := server.NewHTTPServer(cfg.ServerAddr, handler.Init())
	fmt.Println("Server is listening..." + cfg.ServerAddr)
	err = httpServer.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
