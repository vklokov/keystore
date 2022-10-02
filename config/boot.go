package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vklokov/keystore/db"
)

func environment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("* Error loading .env file, %v", err)
	}
}

func Boot() {
	environment()

	db.Connect()
	// db.Migrate()
}
