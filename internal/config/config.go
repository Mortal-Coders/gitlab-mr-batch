package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Gitlab struct {
		Host  string `yaml:"host"`
		Token string `yaml:"token"`
	} `yaml:"gitlab"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config file error: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("yaml decode error: %w", err)
	}
	return config, nil
}
