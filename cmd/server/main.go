// @file:  main.go
// @date:  2021/11/1

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"learning/docs"
	"learning/internal"
	"learning/internal/config"
	"learning/internal/constant"
	"learning/internal/data"

	"go.uber.org/zap"
)

func main() {
	release := flag.Bool("release", false, "Run in release mode.")
	addr := flag.String("addr", ":49091", "HTTP service address.")
	flag.Parse()

	var logger *zap.Logger
	if *release {
		logger, _ = zap.NewProduction(zap.AddCaller())
	} else {
		logger, _ = zap.NewDevelopment(zap.AddCaller())
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	config.Init("config", "yml", []string{"config"})
	cfg := config.GetConfig()

	data.Init(cfg.Database)

	docs.SwaggerInfo.Title = "Chat App Server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is the chat app server."
	docs.SwaggerInfo.Schemes = []string{"http"}

	app := internal.NewApp()
	go func() {
		if err := app.Listen(*addr); err != nil {
			zap.S().Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig
	fmt.Println("Shutting down...")
	go app.Shutdown()
	<-time.After(constant.ShutdownTime)
	fmt.Println("Running cleanup tasks...")
	// cleanup tasks go here
	fmt.Println("Chat server was successful shutdown.")
}
