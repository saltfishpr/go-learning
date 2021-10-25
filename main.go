package main

import (
	"context"
	"time"

	"learning/config"
	"learning/logger"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func init() {
	config.Init()
	logger.Init()
	logger.Info("build date: ", config.BuildDate)
}

func main() {
	defer logger.Sync()

	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   viper.GetStringSlice("etcd.endpoints"),
			DialTimeout: 5 * time.Second,
		},
	)
	if err != nil {
		logger.Fatal("create etcd client error: ", err)
	}
	defer cli.Close()

	putKey(cli, "foo", "bar")
	getKey(cli, "foo")
}

func putKey(cli *clientv3.Client, key string, val string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("etcd.timeout"))*time.Second)
	_, err := cli.Put(ctx, key, val)
	cancel()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Put [%s: %s]", key, val)
}

func getKey(cli *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("etcd.timeout"))*time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		logger.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		logger.Infof("Get [%s: %s]", ev.Key, ev.Value)
	}
}
