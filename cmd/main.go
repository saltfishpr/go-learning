package main

import (
	"learning/internal/config"

	"go.uber.org/zap"
)

var buildTag = "undef"

func main() {
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	zap.S().Info("Build tag: ", buildTag)

	config.Init("config", "yml", []string{"."})
	cfg := config.GetConfig()
	zap.S().Infof("Hello, %s!", cfg.Name)
}
