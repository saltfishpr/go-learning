// @file: main.go
// @date: 2021/10/28

package main

import (
	"time"

	"learning/config"
	"learning/internal/model"
	"learning/logger"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
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

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   viper.GetStringSlice("kafka.brokers"),
			Marshaler: kafka.DefaultMarshaler{},
		},
		nil,
	)
	if err != nil {
		logger.Fatal(err)
	}

	for {
		file := model.File{
			Name:    "file name",
			Level:   1,
			Content: "this is a test file",
		}
		data, _ := json.Marshal(&file)
		msg := message.NewMessage(watermill.NewUUID(), data)
		if err := publisher.Publish(viper.GetString("watermill.topic"), msg); err != nil {
			logger.Fatal("publish message error:", err)
		}
		logger.Infof("publish message: %s, payload: %s", msg.UUID, string(msg.Payload))
		time.Sleep(time.Second)
	}
}
