package main

import (
	"log"
	"ozonTech/muhtarov/internal/app"
)

func main() {
	log.Println("Starting the app...")
	app.Run("configs/config.json")
}
