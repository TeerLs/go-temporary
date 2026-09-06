package main

import (
	"net/http"
	"temporary/config"
	"temporary/internal/product"
	"temporary/pkg/db"
)

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
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("Failed to start the server")
	}
}