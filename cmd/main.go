package main

import (
	"flag"

	"learning/internal/config"
	"learning/internal/logger"
)

var buildTag = "undef"

func main() {
	release := flag.Bool("release", false, "Run in release mode.")
	flag.Parse()

	logger.Init(*release)
	defer logger.Sync()
	logger.Info("build tag: ", buildTag)

	config.Init("config", "yml", []string{"."})
	cfg := config.GetConfig()
	logger.Infof("Hello, %s!", cfg.Name)
}
