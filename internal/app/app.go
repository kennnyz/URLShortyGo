package app

import (
	"fmt"
	"log"
	"os"
	"ozonTech/muhtarov/internal/config"
	"ozonTech/muhtarov/internal/delivery/http"
	"ozonTech/muhtarov/internal/repository"
	"ozonTech/muhtarov/internal/server"
	"ozonTech/muhtarov/internal/service"
)

func Run(configPath string) {
	cfg, err := config.ReadConfig(configPath)
	if err != nil {
		panic(err)
	}

	repos, err := repository.NewRepository(cfg)
	if err != nil {
		panic(err)
	}

	urlShorty := service.NewURLShortyService(repos)
	handler := http.NewHandler(urlShorty)

	httpServer := server.NewHTTPServer(cfg.ServerAddr, handler.Init())
	fmt.Println("Server is listening..." + cfg.ServerAddr)
	err = httpServer.Run()
	if err != nil {
		_, err := os.Stderr.WriteString(err.Error())
		if err != nil {
			log.Println(err)
			return
		}
	}
}
