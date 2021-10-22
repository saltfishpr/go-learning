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

func main() {
	err := config.Init()
	if err != nil {
		logger.Fatal("init config failed: ", err)
	}

	listen, err := net.Listen("tcp", viper.GetString("rpcAddr"))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterHelloServer(grpcServer, &service.HelloServerImpl{})
	logger.Info("server listening at ", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
