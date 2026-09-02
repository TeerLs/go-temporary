package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string `validate:"required"`
}

type ServerConfig struct {
	Address string `validate:"required"`
}

type Config struct {
	DB DBConfig `validate:"required"`
	Server ServerConfig `validate:"required"`
}

func (c *Config) LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		DB: DBConfig{
			DSN: os.Getenv("DSN_DB"),
		},
		Server: ServerConfig{
			Address: os.Getenv("ADDRESS"),
		},
	}

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		log.Fatal("Invalid configuration: ", err)
	}

	return &config
}