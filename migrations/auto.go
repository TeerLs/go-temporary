package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"temporary/internal/product"
)

func main() {

		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&product.Product{})
}