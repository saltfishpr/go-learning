// @file: hello.go
// @date: 2021/10/23

// Package service
package service

import (
	"context"
	"strings"
	"time"

	v1 "learning/api/hello/v1"
	"learning/logger"

	"github.com/moby/moby/pkg/pubsub"
)

// PubSubServerImpl 实现 HelloServer 服务
type PubSubServerImpl struct {
	pub *pubsub.Publisher
	v1.UnimplementedPubSubServer
}

func NewPubSubServerImpl() *PubSubServerImpl {
	return &PubSubServerImpl{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubSubServerImpl) Publish(_ context.Context, arg *v1.String) (*v1.String, error) {
	logger.Info("get publish: ", arg.GetValue())
	p.pub.Publish(arg.GetValue())
	return &v1.String{}, nil
}

func (p *PubSubServerImpl) Subscribe(
	arg *v1.String, stream v1.PubSub_SubscribeServer,
) error {
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
