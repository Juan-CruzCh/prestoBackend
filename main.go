package main

import (
	"log"
	"os"
	"prestoBackend/src"
	"prestoBackend/src/core/config"

	"github.com/joho/godotenv"
)

func main() {
	config.ConfiguracionLog()
	defer config.CerrarLog()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	urlMongo := os.Getenv("URL_MONGO")
	port := os.Getenv("PORT")
	app := src.NewApp(urlMongo)
	app.Run(port)
}
