package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userv1 "github.com/saltfishpr/go-learning/gen/go/user/v1"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":9001", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := userv1.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.CreateUser(ctx, &userv1.CreateUserRequest{
		User: &userv1.CreateUserRequest_User{
			Username: "hello",
			Password: "world!",
			Email:    "ttt@example.com",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
