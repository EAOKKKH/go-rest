package main

import (
	"go-rest/internal/config"
	"go-rest/internal/logger"
	"go-rest/internal/postgres"
	"go-rest/internal/server"
	"os"
)

func main() {
	os.Setenv("CONFIG_PATH", "config/local.yaml")

	cfg := config.LoadConfig()

	logger := logger.NewLogger(cfg)

	dbConn, err := postgres.NewPgConn(cfg)
	if err != nil {
		logger.Fatalf("Error connection to db: %v", err)
	}
	defer dbConn.Close()
	s := server.NewServer(cfg, dbConn, logger)
	s.Run()
}
