package iam

import (
	"MydroX/project-v/pkg/config"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env  string `yaml:"env"`
	Port string `yaml:"port"`
}

func LoadConfig(serviceName string) (*Config, error) {
	f, err := config.Read(serviceName)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
