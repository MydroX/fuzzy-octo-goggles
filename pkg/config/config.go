package config

import (
	"fmt"
	"os"
)

func Read(serviceName string) ([]byte, error) {
	file, err := os.ReadFile(fmt.Sprintf("cmd/%s/config.yml", serviceName))
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}

	return file, nil
}
