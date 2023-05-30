package middleware

import (
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

// Ref: https://github.com/apache/skywalking/blob/master/oap-server/server-starter/src/main/resources/component-libraries.yml
const componentIDGoMQSubscriber = 5013

type SkyWalking struct {
	tracer        *go2sky.Tracer
	operationName string
	logger        *zap.Logger
}

func NewSkyWalking(
	tracer *go2sky.Tracer,
	operationName string,
	logger *zap.Logger,
) *SkyWalking {
	return &SkyWalking{
		tracer:        tracer,
		operationName: operationName,
		logger:        logger,
	}
}

func (sw *SkyWalking) Middleware(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) (messages []*message.Message, err error) {
		span, ctx, err := sw.tracer.CreateEntrySpan(
			msg.Context(),
			sw.operationName,
			func(headerKey string) (string, error) {
				return msg.Metadata[headerKey], nil
			},
		)
		if err != nil {
			sw.logger.Error("create entry span error", zap.Error(err))
			return h(msg)
		}

		msg.SetContext(ctx)
		defer span.End()

		span.SetComponent(componentIDGoMQSubscriber)
		span.SetSpanLayer(agentv3.SpanLayer_MQ) // 4
		defer func() {
			if err != nil {
				span.Error(time.Now(), err.Error())
			}
		}()

		return h(msg)
	}
}
