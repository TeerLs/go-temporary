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
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.WarnLevel)
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