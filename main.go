package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"go.uber.org/zap"

	"learning/decorator"
	_middleware "learning/middleware"
)

func main() {
	reporter, _ := reporter.NewLogReporter()
	tracer, err := go2sky.NewTracer("localhost:30001", go2sky.WithReporter(reporter))
	if err != nil {
		panic(err)
	}
	go2sky.SetGlobalTracer(tracer)

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)

	router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}

	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,

		// The handler function is retried if it returns an error.
		// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          watermill.NewStdLogger(false, false),
		}.Middleware,

		// Recoverer handles panics from handlers.
		// In this case, it passes them as errors to the Retry middleware.
		middleware.Recoverer,

		_middleware.SkyWalking(tracer, "consumer111", "localhost", logger),
	)

	// For simplicity, we are using the gochannel Pub/Sub here,
	// You can replace it with any Pub/Sub implementation, it will work the same.
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{
			// Subscribe 时返回的通道缓冲区大小
			OutputChannelBuffer: 10,
			// 是否持久化消息(保存在slice中), 一般不要设置为 true, 会导致内存持续上涨
			Persistent: false,
			// 阻塞推送消息直到有消费者消费
			BlockPublishUntilSubscriberAck: false,
		},
		watermill.NewStdLogger(false, false),
	)

	// just for debug, we are printing all messages received on `incoming_messages_topic`
	router.AddNoPublisherHandler(
		"print_incoming_messages",
		"incoming_messages_topic",
		pubSub,
		printMessages,
	)

	go func() {
		if err := router.Run(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Wait for consumer running
	<-router.Running()
	// Producing some messages
	go publishMessages(pubSub)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func publishMessages(publisher message.Publisher) {
	publisher, _ = decorator.SkyWalkingPublisherDecorator(
		go2sky.GetGlobalTracer(),
		"greet_event",
		"localhost",
		zap.L(),
	)(publisher)

	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
		middleware.SetCorrelationID(watermill.NewUUID(), msg)
		span, ctx, err := go2sky.GetGlobalTracer().CreateEntrySpan(
			msg.Context(),
			"provider111",
			func(headerKey string) (string, error) {
				return msg.Metadata[headerKey], nil
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		span.SetComponent(5000)
		span.SetSpanLayer(4) // 4
		msg.SetContext(ctx)

		log.Printf("sending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))

		if err := publisher.Publish("incoming_messages_topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}
