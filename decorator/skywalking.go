package decorator

import (
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

// Ref: https://github.com/apache/skywalking/blob/master/oap-server/server-starter/src/main/resources/component-libraries.yml
const componentIDUnknown = 0

type skyWalkingPublisherDecorator struct {
	pub message.Publisher

	tracer *go2sky.Tracer

	operationName string
	// localhost if use go channel, kafka addrs if use kafka
	peer string

	logger *zap.Logger
}

func SkyWalkingPublisherDecorator(
	tracer *go2sky.Tracer,
	operationName string,
	peer string,
	logger *zap.Logger,
) message.PublisherDecorator {
	return func(pub message.Publisher) (message.Publisher, error) {
		return &skyWalkingPublisherDecorator{
			pub:           pub,
			tracer:        tracer,
			operationName: operationName,
			peer:          peer,
			logger:        logger,
		}, nil
	}
}

func (d *skyWalkingPublisherDecorator) Publish(
	topic string,
	messages ...*message.Message,
) (err error) {
	if len(messages) == 0 {
		return d.pub.Publish(topic)
	}

	span, err := d.tracer.CreateExitSpan(
		messages[0].Context(),
		d.operationName,
		d.peer,
		func(headerKey, headerValue string) error {
			for _, msg := range messages {
				msg.Metadata[headerKey] = headerValue
			}
			return nil
		},
	)
	if err != nil {
		d.logger.Error("create exit span error", zap.Error(err))
		return d.pub.Publish(topic, messages...)
	}

	defer span.End()

	span.SetComponent(componentIDUnknown)
	span.SetSpanLayer(agentv3.SpanLayer_MQ)
	span.Tag(go2sky.TagMQTopic, topic)

	defer func() {
		if err != nil {
			span.Error(time.Now(), err.Error())
		}
	}()

	return d.pub.Publish(topic, messages...)
}

func (d *skyWalkingPublisherDecorator) Close() error {
	return d.pub.Close()
}
