// @file: hello.go
// @date: 2021/10/23

// Package service 实现 hello.v1 中定义的服务
package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	v1 "learning/api/hello/v1"
	"learning/logger"

	"github.com/moby/moby/pkg/pubsub"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// PubSubServerImpl 实现 v1.PubSubServer
type PubSubServerImpl struct {
	v1.UnimplementedPubSubServer
	pub  *pubsub.Publisher
	auth Authentication
}

func NewPubSubServerImpl() *PubSubServerImpl {
	return &PubSubServerImpl{
		auth: Authentication{
			Username: viper.GetString("auth.username"),
			Password: viper.GetString("auth.password"),
		},
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubSubServerImpl) Publish(ctx context.Context, arg *v1.String) (*v1.String, error) {
	if err := p.auth.auth(ctx); err != nil {
		return nil, err
	}
	logger.Info("get publish: ", arg.GetValue())
	p.pub.Publish(arg.GetValue())
	return &v1.String{}, nil
}

func (p *PubSubServerImpl) Subscribe(arg *v1.String, stream v1.PubSub_SubscribeServer) error {
	if err := p.auth.auth(stream.Context()); err != nil {
		return err
	}
	ch := p.pub.SubscribeTopic(
		func(v interface{}) bool {
			if key, ok := v.(string); ok {
				if strings.HasPrefix(key, arg.GetValue()) {
					return true
				}
			}
			return false
		},
	)

	for v := range ch {
		if err := stream.Send(&v1.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

type Authentication struct {
	Username string
	Password string
}

func (a Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"username": a.Username, "password": a.Password}, nil
}

func (a Authentication) RequireTransportSecurity() bool {
	return false
}

// auth 从 context 中获取认证信息并比对
func (a Authentication) auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid, appkey string
	if val, ok := md["username"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}
	if appid != a.Username || appkey != a.Password {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}
