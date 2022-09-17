package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{"localhost:2379"},
			DialTimeout: 5 * time.Second,
		},
	)
	if err != nil {
		log.Fatal("create etcd client error: ", err)
	}
	defer cli.Close()

	putKey(cli, "foo", "bar")
	getKey(cli, "foo")
}

func putKey(cli *clientv3.Client, key string, val string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	_, err := cli.Put(ctx, key, val)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Put [%s: %s]", key, val)
}

func getKey(cli *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		log.Printf("Get [%s: %s]", ev.Key, ev.Value)
	}
}
