package db

import (
	"fmt"
	"os"

	"github.com/vklokov/keystore/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() error {
	credentials := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Berlin",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	conn, err := gorm.Open(postgres.Open(credentials), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Conn = conn

	return nil
}

func Migrate() {
	Conn.AutoMigrate(&entities.User{})
}
