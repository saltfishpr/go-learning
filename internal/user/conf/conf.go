// Package conf .
package conf

type Config struct {
	Port     int    `json:"port"`
	LogLevel string `json:"log_level"`
	OTEL     OTEL   `json:"otel"`
}

type OTEL struct {
	ServiceName   string `json:"service_name"`
	CollectorAddr string `json:"collector_addr"`
}

func New() *Config {
	return &Config{
		Port:     8080,
		LogLevel: "info",
	}
}
