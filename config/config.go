package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string
}

type ServerConfig struct {
	Address string
}

type Config struct {
	DB DBConfig
	Server ServerConfig
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

	return &config
}