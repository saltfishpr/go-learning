// @file: hello.go
// @date: 2021/10/23

// Package service
package service

import (
	v1 "learning/api/hello/v1"
	"learning/logger"
)

// HelloService 定义rpc调用接口
type HelloService interface {
	Hello(*v1.String, *v1.String) error
}

// HelloServiceImpl 实现 HelloService 服务
type HelloServiceImpl struct{}

// Hello 返回 "hello {{request}}"
func (p *HelloServiceImpl) Hello(request *v1.String, reply *v1.String) error {
	logger.Infof("get request: %s", request.GetValue())
	reply.Value = "hello:" + request.GetValue()
	return nil
}
