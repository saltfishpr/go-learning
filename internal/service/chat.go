// @description:
// @file: chat.go
// @date: 2021/11/22

package service

import (
	"context"
	"errors"
	"fmt"

	"learning/internal/common/connstorage"
	"learning/internal/data"
	"learning/internal/model"
	"learning/internal/utils"
	"learning/logger"
)

func ProcessMessage(ctx context.Context, data []byte) error {
	account := ctx.Value("account").(string)
	message := new(model.RecvMessage)
	err := utils.JsonUnmarshal(data, message)
	if err != nil {
		return err
	}
	switch *message.Mode {
	case model.ToUser:
		return sendToUser(account, message)
	case model.ToHub:
		return sendToHub(account, message)
	default:
		return errors.New("wrong message mode")
	}
}

func sendToUser(account string, message *model.RecvMessage) error {
	return send(*message.ID, *message.Content)
}

func sendToHub(account string, message *model.RecvMessage) error {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return err
	}
	hubEntity, err := data.GetHubByHID(*message.ID)
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
		if err := send(*userEntity.Account, *message.Content); err != nil {
			logger.Error(err)
		}
	}
	return nil
}

func send(account string, message string) error {
	sendMessage := new(model.SendMessage)
	sendMessage.Content = message
	if connection, ok := connstorage.Get(account); ok {
		return connection.WriteJSON(sendMessage)
	}
	return fmt.Errorf("%s is offline", account)
}
