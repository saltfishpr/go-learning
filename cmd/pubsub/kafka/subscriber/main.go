// @file: main.go
// @date: 2021/10/28

package main

import (
	"context"

	"learning/config"
	"learning/internal/model"
	"learning/logger"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	config.Init()
	logger.Init()
	logger.Infof("build date: %s", config.BuildDate)
}

func main() {
	defer logger.Sync()

	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               viper.GetStringSlice("kafka.brokers"),
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         viper.GetString("kafka.consumer_group"),
		},
		nil,
	)
	if err != nil {
		logger.Fatal(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), viper.GetString("watermill.topic"))
	if err != nil {
		logger.Fatal(err)
	}

	for msg := range messages {
		var file model.File
		_ = json.Unmarshal(msg.Payload, &file)
		logger.Infof("received message: %s, payload: %+v", msg.UUID, file)
		msg.Ack()
	}
}
