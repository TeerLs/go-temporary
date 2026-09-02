package db

import (
	"temporary/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(db *config.DBConfig) *DB {
	database, err := gorm.Open(postgres.Open(db.DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	return &DB{DB: database}
}
