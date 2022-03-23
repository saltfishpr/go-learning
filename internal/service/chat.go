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

type user struct {
	account string
}

func (u user) auth(s string) bool {
	// TODO implement me
	panic("implement me")
}

func (u user) send(message *model.RecvMessage) error {
	// TODO implement me
	panic("implement me")
}

type hub struct {
	hid string
}

func (h hub) auth(s string) bool {
	// TODO implement me
	panic("implement me")
}

func (h hub) send(message *model.RecvMessage) error {
	// TODO implement me
	panic("implement me")
}

func sendToUser(account string, message *model.RecvMessage) error {
	if connection, ok := connstorage.Get(account); ok {
		return connection.WriteJSON(message)
	}
	return fmt.Errorf("%s is offline", account)
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
		account := topicStr[constant.TopicPrefixLen:]
		return user{account: account}
	case "hub":
		hid := topicStr[constant.TopicPrefixLen:]
		return hub{hid: hid}
	default:
		return nil
	}
}

func GetMessagesPagination(
	account string,
	query *model.MessagesPaginationRequest,
) (*model.MessagesPaginationResponse, error) {
	return nil, nil
}
