package main

import (
	"net/http"
	"os"
	"temporary/config"
	"temporary/internal/auth"
	"temporary/internal/product"
	"temporary/pkg/db"
	"temporary/pkg/jwt"
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

	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config:       cfg.Auth,
		SessionsStore: auth.NewSessionStore(),
		JWT:          jwt.NewJWT(cfg.Auth.SecretKey),
	})

	product.NewProductHandler(router, &product.ProductHandlerDeps{
		DB: database,
	})

	server := http.Server{
		Addr:    cfg.Server.Address,
		Handler: middleware.AuthMiddleware(jwt.NewJWT(cfg.Auth.SecretKey))(middleware.LoggingMiddleware(router)),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("Failed to start the server")
	}
}