package main

import (
	"net/http"
	"temporary/config"
	"temporary/pkg/db"
)

func main() {
	cfg := config.Config{}
	cfg = *cfg.LoadConfig()
	_ = db.NewDB(&cfg.DB)

	server := http.Server{
		Addr:    cfg.Server.Address,
		Handler: nil,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("Failed to start the server")
	}
}