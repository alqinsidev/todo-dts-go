package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env")
	}

	return os.Getenv("MONGO_URI")
}

func EnvPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env")
	}

	return os.Getenv("PORT")
}
