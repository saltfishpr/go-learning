// @description:
// @file: chat.go
// @date: 2021/11/22

package service

import (
	"errors"

	"learning/internal/data"
	"learning/internal/model"
	"learning/internal/utils"
)

func ProcessMessage(account string, data []byte) error {
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
	sendMessage := new(model.SendMessage)
	sendMessage.Content = *message.Content

	connection := utils.GetConnection(*message.ID)
	return connection.WriteJSON(sendMessage)
}

func sendToHub(account string, message *model.RecvMessage) error {
	sendMessage := new(model.SendMessage)
	sendMessage.Content = *message.Content

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
		return errors.New("user has not joined the chat room")
	}

	userEntities, err := data.GetUsersInHub(hubEntity)
	if err != nil {
		return err
	}
	for _, userEntity := range userEntities {
		connection := utils.GetConnection(*userEntity.Account)
		err := connection.WriteJSON(sendMessage)
		if err != nil {
			return err
		}
	}
	return nil
}
