// @file: hello.go
// @date: 2021/10/23

// Package service
package service

import (
	"context"
	"io"

	v1 "learning/api/hello/v1"
	"learning/logger"
)

// HelloServerImpl 实现 HelloServer 服务
type HelloServerImpl struct {
	v1.UnimplementedHelloServer
}

// Hello 返回 "hello {{request}}"
func (s *HelloServerImpl) Hello(_ context.Context, request *v1.String) (*v1.String, error) {
	logger.Infof("get request: %s", request.GetValue())
	reply := &v1.String{Value: "hello " + request.GetValue()}
	return reply, nil
}

func (s *HelloServerImpl) Channel(stream v1.Hello_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		logger.Infof("channel receive: %s", args.GetValue())
		reply := &v1.String{Value: "hello " + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
