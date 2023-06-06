package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	REDIS_PORT     = ""
	REDIS_PASSWORD = ""
	SERVER_PORT    = ""
	REDIS_HOST     = ""
)

func LoadConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load configs", err)
	}

	REDIS_PORT = os.Getenv("REDIS_PORT")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	SERVER_PORT = os.Getenv("SERVER_PORT")
	REDIS_HOST = os.Getenv("REDIS_HOST")
}
