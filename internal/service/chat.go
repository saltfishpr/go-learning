// @description: 消息服务
// @file: chat.go
// @date: 2021/11/22

package service

import (
	"context"
	"errors"
	"fmt"

	"learning/internal/common/connstorage"
	"learning/internal/constant"
	"learning/internal/model"
	"learning/internal/utils"

	"go.uber.org/zap"
)

type topic interface {
	auth(string) bool
	send(*model.RecvMessage) error
}

type userTopic struct {
	username string
}

func (u userTopic) auth(s string) bool {
	// TODO implement me
	panic("implement me")
}

func (u userTopic) send(message *model.RecvMessage) error {
	// TODO implement me
	panic("implement me")
}

type hubTopic struct {
	hid string
}

func (h hubTopic) auth(s string) bool {
	// TODO implement me
	panic("implement me")
}

func (h hubTopic) send(message *model.RecvMessage) error {
	// TODO implement me
	panic("implement me")
}

func sendToUser(username string, message *model.RecvMessage) error {
	if connection, ok := connstorage.Get(username); ok {
		return connection.WriteJSON(message)
	}
	return fmt.Errorf("%s is offline", username)
}

func ProcessMessage(ctx context.Context, data []byte) error {
	message := new(model.RecvMessage)
	err := utils.JsonUnmarshal(data, message)
	if err != nil {
		return err
	}
	zap.S().Info(message)
	t := getTopic(message.Topic)
	if t == nil {
		return errors.New("wrong topic")
	}
	// TODO: Save message to database
	return t.send(message)
}

func getTopic(topicStr string) topic {
	switch topicStr[:constant.TopicPrefixLen] {
	case "usr":
		username := topicStr[constant.TopicPrefixLen:]
		return userTopic{username: username}
	case "hub":
		hid := topicStr[constant.TopicPrefixLen:]
		return hubTopic{hid: hid}
	default:
		return nil
	}
}

func GetMessagesPagination(
	username string,
	query *model.MessagesPaginationRequest,
) (*model.MessagesPaginationResponse, error) {
	return nil, nil
}
