// @file:  main.go
// @date:  2021/11/1

package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"learning/internal"
	"learning/internal/config"
	"learning/internal/constant"
	"learning/internal/log"
)

func main() {
	logger := log.New()
	defer logger.Sync()

	addr := flag.String("addr", ":49091", "HTTP service address.")
	flag.Parse()

	config.Init("config", "yml", []string{"config"})

	app := internal.NewApp(logger)
	go func() {
		if err := app.Listen(*addr); err != nil {
			logger.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig // wait for signal

	logger.Info("Shutting down ...")
	go func() {
		err := app.Shutdown()
		if err != nil {
			logger.Fatal(err)
		}
	}()
	<-time.After(constant.ShutdownTime)
	logger.Info("Running cleanup tasks...")
	// cleanup tasks go here
	logger.Info("Chat server was successful shutdown.")
}
