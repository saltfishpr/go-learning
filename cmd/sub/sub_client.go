// @file: sub_client.go
// @date: 2021/10/24

package main

import (
	"context"
	"io"

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
}

func main() {
	defer logger.Sync()
	logger.Info("build date:", config.BuildDate)

	auth := service.Authentication{
		Username: viper.GetString("auth.username"),
		Password: viper.GetString("auth.password"),
	}
	conn, err := grpc.Dial(viper.GetString("rpcAddr"), grpc.WithInsecure(), grpc.WithPerRPCCredentials(auth))
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewPubSubClient(conn)

	stream, err := client.Subscribe(context.Background(), &v1.String{Value: "golang:"})
	if err != nil {
		logger.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Fatal(err)
		}

		logger.Debug(reply.GetValue())
	}
}
