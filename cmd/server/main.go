// @file: main.go
// @date: 2021/11/1

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
	go func() {
		if err := app.Listen(*addr); err != nil {
			logger.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	_ = <-sig
	fmt.Println("Shutting down...")
	_ = app.Shutdown()
	fmt.Println("Running cleanup tasks...")
	// cleanup tasks go here
	fmt.Println("Chat server was successful shutdown.")
}
