package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// load config
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}
}
