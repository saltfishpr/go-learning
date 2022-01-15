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

	"learning/config"
	"learning/docs"
	"learning/internal"
	"learning/internal/data"
	"learning/internal/logger"
)

var buildTag = "undef"

func main() {
	release := flag.Bool("release", false, "Run in release mode.")
	addr := flag.String("addr", ":49091", "HTTP service address.")
	flag.Parse()

	logger.Init(*release)
	defer logger.Sync()
	logger.Info("build tag: ", buildTag)

	config.Init("config", "yml", []string{"config"})

	data.Init(config.GetString("database"))

	docs.SwaggerInfo.Title = "Chat App Server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is the chat app server."
	docs.SwaggerInfo.Schemes = []string{"http"}

	app := internal.NewApp()
	go func() {
		if err := app.Listen(*addr); err != nil {
			logger.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig
	fmt.Println("Shutting down...")
	go app.Shutdown()
	<-time.After(config.ShutdownTime)
	fmt.Println("Running cleanup tasks...")
	// cleanup tasks go here
	fmt.Println("Chat server was successful shutdown.")
}
