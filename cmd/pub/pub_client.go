// @file: pub_client.go
// @date: 2021/10/24

package main

import (
	"context"

	v1 "learning/api/hello/v1"
	"learning/config"
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

	conn, err := grpc.Dial(viper.GetString("rpcAddr"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewPubSubClient(conn)

	_, err = client.Publish(context.Background(), &v1.String{Value: "golang: hello Go"})
	if err != nil {
		logger.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &v1.String{Value: "docker: hello Docker"})
	if err != nil {
		logger.Fatal(err)
	}
}
