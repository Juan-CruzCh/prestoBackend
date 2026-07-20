package config

import "os"

type configEnv struct {
	MongoURI string
	Database string
	Port     string
}

var AppEnv configEnv

func LoadConfig() {
	AppEnv = configEnv{
		MongoURI: os.Getenv("MONGO_URI"),
		Database: os.Getenv("DATABASE"),
		Port:     os.Getenv("PORT"),
	}
}
