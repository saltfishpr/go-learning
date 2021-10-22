// @file: client.go
// @date: 2021/10/23

package main

import (
	"context"

	v1 "learning/api/hello/v1"
	"learning/config"
	"learning/logger"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	err := config.Init()
	if err != nil {
		logger.Fatal("init config failed: ", err)
	}

	conn, err := grpc.Dial(viper.GetString("rpcAddr"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewHelloClient(conn)
	reply, err := client.Hello(context.Background(), &v1.String{Value: "world"})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(reply.GetValue())
}
