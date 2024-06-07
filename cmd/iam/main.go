package main

import (
	"MydroX/project-v/internal/iam"
	"MydroX/project-v/internal/iam/rest"
	"MydroX/project-v/pkg/logger"
	"log"
)

func main() {
	cfg, err := iam.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	logger := logger.New(cfg.Env)

	logger.Info("starting server...")
	rest.NewServer(cfg, *logger)
}
