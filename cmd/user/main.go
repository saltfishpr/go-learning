package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/oklog/run"
	"github.com/samber/do"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/saltfishpr/go-learning/internal/user"
	"github.com/saltfishpr/go-learning/internal/user/conf"
	"github.com/saltfishpr/go-learning/internal/user/server"
)

var cfgFile string

func init() {
	flag.StringVar(&cfgFile, "config", "configs/config.yaml", "config file path")
}

func main() {
	flag.Parse()

	injector, err := user.NewInjector(cfgFile)
	if err != nil {
		panic(err)
	}

	var g run.Group

	httpServer := server.NewHTTP(injector)
	grpcServer := server.NewGRPC(injector)

	logger := do.MustInvoke[*zap.Logger](injector)
	config := do.MustInvoke[*conf.Config](injector)

	logger.Info(
		"start to listen",
		zap.Int("port", config.Port),
		zap.String("log_level", logger.Level().String()),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	g.Add(func() error {
		h2Handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				httpServer.ServeHTTP(w, r)
			}
		}), &http2.Server{})
		return http.Serve(lis, h2Handler)
	}, func(err error) {
		lis.Close()
	})

	go func() {
		if err := g.Run(); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	grpcServer.GracefulStop()
}
