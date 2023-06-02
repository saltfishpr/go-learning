package middleware

import (
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

// Ref: https://github.com/apache/skywalking/blob/master/oap-server/server-starter/src/main/resources/component-libraries.yml
const componentIDUnknown = 0

func SkyWalking(
	tracer *go2sky.Tracer,
	operationName string,
	peer string,
	logger *zap.Logger,
) message.HandlerMiddleware {
	return func(h message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) ([]*message.Message, error) {
			span, ctx, err := tracer.CreateEntrySpan(msg.Context(), operationName, func(headerKey string) (string, error) {
				return msg.Metadata[headerKey], nil
			})
			if err != nil {
				logger.Error("create entry span error", zap.Error(err))
				return h(msg)
			}

			msg.SetContext(ctx)
			defer span.End()

			span.SetComponent(componentIDUnknown)
			span.SetPeer(peer)
			span.SetSpanLayer(agentv3.SpanLayer_MQ)
			span.Tag(go2sky.TagMQTopic, message.SubscribeTopicFromCtx(ctx))

			defer func() {
				if err != nil {
					span.Error(time.Now(), err.Error())
				}
			}()

			return h(msg)
		}
	}
}
