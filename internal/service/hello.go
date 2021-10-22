// @file: hello.go
// @date: 2021/10/23

// Package service
package service

import (
	"context"

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
	reply := &v1.String{}
	reply.Value = "hello " + request.GetValue()
	return reply, nil
}
