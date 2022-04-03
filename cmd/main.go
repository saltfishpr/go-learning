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
	"learning/internal/data"
	"learning/internal/log"
)

func main() {
	logger := log.New()
	defer logger.Sync()

	addr := flag.String("addr", ":49091", "HTTP service address.")
	flag.Parse()

	config.Init("config", "yml", []string{"."})

	connection, err := data.NewPostgres()
	if err != nil {
		logger.Fatal("connect to database error: ", err)
	}
	if ok, err := connection.IsConnected(); !ok || err != nil {
		logger.Fatal("ping database error: ", err)
	}

	app := internal.NewApp(logger, connection)
	go func() {
		if err := app.Listen(*addr); err != nil {
			logger.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig // wait for signal
	logger.Info("Shutting down ...")

	quit := make(chan struct{})
	go func() {
		err := app.Shutdown()
		if err != nil {
			logger.Fatal(err)
		}
		quit <- struct{}{}
	}()

	select {
	case <-time.After(constant.ShutdownTimeout):
		logger.Warn("Force shutdown with timeout")
	case <-quit:
	}
	logger.Info("Running cleanup tasks...")
	// cleanup tasks go here
	logger.Info("Chat server was successful shutdown.")
}
