package main

import (
	"net/http"
	"os"
	"temporary/config"
	"temporary/internal/product"
	"temporary/pkg/db"
	"temporary/pkg/middleware"

	log "github.com/sirupsen/logrus"
)

func init() {
  log.SetFormatter(&log.JSONFormatter{})

  log.SetOutput(os.Stdout)

  log.SetLevel(log.InfoLevel)
}

func main() {
	cfg := config.Config{}
	cfg = *cfg.LoadConfig()
	database := db.NewDB(&cfg.DB)

	router := http.NewServeMux()

	product.NewProductHandler(router, &product.ProductHandlerDeps{
		DB: database,
	})

	server := http.Server{
		Addr:    cfg.Server.Address,
		Handler: middleware.LoggingMiddleware(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("Failed to start the server")
	}
}