package decorator

import (
	"errors"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

// Ref: https://github.com/apache/skywalking/blob/master/oap-server/server-starter/src/main/resources/component-libraries.yml
const componentIDGoMQPublisher = 5013

type PublisherSkyWalkingDecorator struct {
	pub    message.Publisher
	tracer *go2sky.Tracer

	operationName string
	// localhost if use go channel, kafka addrs if use kafka
	peer string

	logger *zap.Logger
}

func NewPublisherSkyWalkingDecorator(
	pub message.Publisher,
	operationName string,
	peer string,
	tracer *go2sky.Tracer,
	logger *zap.Logger,
) *PublisherSkyWalkingDecorator {
	return &PublisherSkyWalkingDecorator{
		pub:           pub,
		tracer:        tracer,
		operationName: operationName,
		peer:          peer,
		logger:        logger,
	}
}

func (d *PublisherSkyWalkingDecorator) Publish(
	topic string,
	messages ...*message.Message,
) (err error) {
	if len(messages) == 0 {
		return d.pub.Publish(topic)
	}

	var errs []error
	for _, msg := range messages {
		errs = append(errs, d.publish(topic, msg))
	}

	return errors.Join(errs...)
}

func (d *PublisherSkyWalkingDecorator) publish(topic string, msg *message.Message) (err error) {
	span, err := d.tracer.CreateExitSpan(
		msg.Context(),
		d.operationName,
		d.peer,
		func(headerKey, headerValue string) error {
			msg.Metadata[headerKey] = headerValue
			return nil
		},
	)
	if err != nil {
		d.logger.Error("create exit span error", zap.Error(err))
		return d.pub.Publish(topic, msg)
	}

	defer span.End()

	span.SetComponent(componentIDGoMQPublisher)
	span.SetSpanLayer(agentv3.SpanLayer_MQ) // 4
	defer func() {
		if err != nil {
			span.Error(time.Now(), err.Error())
		}
	}()

	return d.pub.Publish(topic, msg)
}

func (d *PublisherSkyWalkingDecorator) Close() error {
	return d.pub.Close()
}
