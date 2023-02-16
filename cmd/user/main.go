package main

import (
	"flag"
	"net"
	"os"
	"os/signal"

	"github.com/samber/do"
	"go.uber.org/zap"

	"github.com/saltfishpr/go-learning/internal/user"
	"github.com/saltfishpr/go-learning/internal/user/conf"
	"github.com/saltfishpr/go-learning/internal/user/server"
)

var (
	cfgFile string
	release bool
)

func init() {
	flag.StringVar(&cfgFile, "config", "configs/config.yaml", "config file path")
	flag.BoolVar(&release, "release", false, "release mode")
}

func main() {
	flag.Parse()

	injector, err := user.NewInjector(cfgFile, release)
	if err != nil {
		panic(err)
	}

	s := server.NewGRPCServer(injector)

	logger := do.MustInvoke[*zap.Logger](injector)
	config := do.MustInvoke[*conf.Config](injector)

	lis, err := net.Listen("tcp", config.Addr)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	if err := s.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	s.GracefulStop()
}
