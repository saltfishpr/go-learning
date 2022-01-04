// @description:
// @file: chat.go
// @date: 2021/11/22

package service

import (
	"context"
	"errors"
	"fmt"

	"learning/config"
	"learning/internal/common/connstorage"
	"learning/internal/data"
	"learning/internal/model"
	"learning/internal/utils"
	"learning/logger"
)

type Topic interface {
	auth(string) bool
	send(*model.RecvMessage) error
}

type user struct {
	account string
}

func (u user) auth(account string) bool {
	return true
}

func (u user) send(message *model.RecvMessage) error {
	return sendToUser(u.account, message)
}

type hub struct {
	hid string
}

func (h hub) auth(account string) bool {
	return true
}

func (h hub) send(message *model.RecvMessage) error {
	userEntity, err := data.GetUserByAccount(message.From)
	if err != nil {
		return err
	}
	hubEntity, err := data.GetHubByHID(h.hid)
	if err != nil {
		return err
	}
	ok, err := data.IsInHub(hubEntity, userEntity)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user has not joined the hub")
	}
	userEntities, err := data.GetUsersInHub(hubEntity)
	if err != nil {
		return err
	}

	for _, userEntity := range userEntities {
		if err := sendToUser(*userEntity.Account, message); err != nil {
			logger.Error(err)
		}
	}
	return nil
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
	logger.Info(message)
	topic := getTopic(message.Topic)
	if topic == nil {
		return errors.New("wrong topic")
	}
	return topic.send(message)
}

func getTopic(topic string) Topic {
	switch topic[:config.TopicPrefixLen] {
	case "usr":
		account := topic[config.TopicPrefixLen:]
		return user{account: account}
	case "hub":
		hid := topic[config.TopicPrefixLen:]
		return hub{hid: hid}
	default:
		return nil
	}
}
