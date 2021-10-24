// @file: client.go
// @date: 2021/10/23

package main

import (
	"context"
	"io"
	"time"

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
	stream, err := client.Channel(context.Background())
	if err != nil {
		logger.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&v1.String{Value: "world"}); err != nil {
				logger.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info(reply.GetValue())
	}
}
