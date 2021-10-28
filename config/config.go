// @file: config.go
// @date: 2021/10/22

// Package config 读取配置
package config

import (
	"github.com/spf13/viper"
)

var (
	BuildDate string
	Release   = "false"
)

// Init 初始化配置
func Init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {
		panic(err)
	}
}
