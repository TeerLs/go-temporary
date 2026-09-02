package main

import (
	"net/http"
	"temporary/config"
	"temporary/pkg/db"
)

func main() {
	config := config.Config{}
	cfg := config.LoadConfig()
	_ = db.NewDB(&cfg.DB)

	server := http.Server{
		Addr:    cfg.Server.Address,
		Handler: nil,
	}

	server.ListenAndServe()
}