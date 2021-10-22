package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"learning/config"
	"learning/internal/service"
	"learning/logger"

	"github.com/spf13/viper"
)

//go:generate protoc --proto_path=. --go_out=.  api/hello/v1/hello.proto

const HelloServiceName = "api"

// RegisterHelloService 注册一个 HelloService rpc服务
func RegisterHelloService(svc service.HelloService) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func rpcService(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("rpc accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	var conn io.ReadWriteCloser = struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: r.Body,
		Writer:     w,
	}
	err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	if err != nil {
		logger.Error("http service error: ")
	}
}

func main() {
	err := config.Init()
	if err != nil {
		logger.Fatal("init config failed: ", err)
	}

	err = RegisterHelloService(new(service.HelloServiceImpl))
	if err != nil {
		logger.Fatal("Register error: ", err)
	}

	rpcAddr := fmt.Sprintf("%s:%d", "", viper.Get("rpcPort"))
	logger.Infof("rpc service addr is %s", rpcAddr)
	listener, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		logger.Fatal("ListenTCP error: ", err)
	}
	go rpcService(listener)

	httpAddr := fmt.Sprintf("%s:%d", "", viper.Get("httpPort"))
	logger.Infof("http service addr is %s", httpAddr)
	http.HandleFunc("/"+HelloServiceName, httpHandler)
	err = http.ListenAndServe(httpAddr, nil)
	if err != nil {
		logger.Fatal("http ListenAndServe error: ", err)
	}
}
