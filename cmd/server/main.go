// @file: main.go
// @date: 2021/11/1

package main

import (
	"flag"

	"learning/config"
	"learning/internal"
	"learning/logger"
)

var BuildDate string

func init() {
	config.Init("config", "yml", []string{"config"})
	logger.Init()
	logger.Info("build date: ", BuildDate)
}

func main() {
	defer logger.Sync()
	addr := flag.String("addr", ":49091", "http service address")
	flag.Parse()

	logger.Info("Listening: ", *addr)
	app := internal.NewApp()
	logger.Fatal(app.Listen(*addr))
}
