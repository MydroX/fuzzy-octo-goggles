package main

import (
	"MydroX/project-v/internal/iam"
	"MydroX/project-v/pkg/logger"
	"log"
)

const serviceName = "iam"

func main() {
	cfg, err := iam.LoadConfig(serviceName)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	logger := logger.New(cfg.Env)

	logger.Info("starting server...")
	iam.NewServer(cfg, *logger)
}
