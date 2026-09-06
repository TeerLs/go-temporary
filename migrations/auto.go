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

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_DB")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		panic("Failed to migrate the database")
	}
}