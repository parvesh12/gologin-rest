package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	err := godotenv.Load(".env")

	if err != nil {

		fmt.Printf("Error load env: %s ", err)

	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")) //Build connection string

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {

		panic(err)
	}

	return db
}
