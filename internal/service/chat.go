// @description:
// @file: chat.go
// @date: 2021/11/22

package service

import (
	"learning/internal/model"
	"learning/internal/utils"
	"learning/logger"
)

func ProcessMessage(data []byte) {
	message := new(model.Message)
	err := utils.JsonUnmarshal(data, message)
	if err != nil {
		logger.Error("process message error: ", err)
		return
	}
	switch *message.Mode {
	case model.ToUser:
		sendToUser()
	case model.ToHub:
		sendToHub()
	default:
		logger.Error("process message error: ", err)
	}
}

func sendToUser() {

}

func sendToHub() {

}
