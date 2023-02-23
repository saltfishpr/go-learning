// Package conf .
package conf

type Config struct {
	Port      int    `json:"port"`
	SecretKey string `json:"secret_key"`
}
