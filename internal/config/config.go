// @file: config.go
// @date: 2022/3/21

package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config represents the main config for the application.
type Config struct {
	NodeID   int64  `mapstructure:"node_id"`
	Database string `mapstructure:"database"`
}

// Init loads config.
func Init(fileName string, fileExtension string, configPaths []string) {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileExtension)
	for _, configPath := range configPaths {
		viper.AddConfigPath(configPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			zap.S().Fatal(err)
		}
	}
}

// GetConfig returns the users' config.
func GetConfig() (config Config) {
	if err := viper.Unmarshal(&config); err != nil {
		zap.S().Fatal("Error parsing config", err)
	}
	zap.S().Debug("Config: ", config)
	return
}
