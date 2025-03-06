package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr      string           `yaml:"addr"`
	Providers []ProviderConfig `yaml:"providers"`
}

type ProviderConfig struct {
	Name      string `yaml:"name"`
	BaseURL   string `yaml:"baseURL"`
	APIKey    string `yaml:"apiKey"`
	TimeoutMs int    `yaml:"timeoutMs"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
