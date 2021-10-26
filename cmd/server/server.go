// @file: server.go
// @date: 2021/10/23

package main

import (
	"net"

	v1 "learning/api/hello/v1"
	"learning/config"
	"learning/internal/service"
	"learning/logger"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	config.Init()
	logger.Init()
	logger.Info("build date:", config.BuildDate)
}

func main() {
	defer logger.Sync()

	grpcServer := grpc.NewServer()
	v1.RegisterPubSubServer(grpcServer, service.NewPubSubServerImpl())

	listen, err := net.Listen("tcp", viper.GetString("rpcAddr"))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	logger.Info("server listening at ", listen.Addr())

	if err := grpcServer.Serve(listen); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
