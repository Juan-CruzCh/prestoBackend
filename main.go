package main

import (
	"log"
	"prestoBackend/src/app/config"
	logApp "prestoBackend/src/app/log"
	"prestoBackend/src/app/server"

	"github.com/joho/godotenv"
)

func main() {
	logApp.ConfiguracionLog()
	defer logApp.CerrarLog()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadConfig()
	app := server.NewApp()
	app.Run()
}
