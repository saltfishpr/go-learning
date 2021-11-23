// @description: 处理用户发送的消息
// @file: chat.go
// @date: 2021/11/22

package v1

import (
	"learning/internal/service"
	"learning/internal/utils"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ChatHandler(c *websocket.Conn) {
	account := utils.GetUserAccountWebsocketConn(c)
	utils.Connect(account, c)
	defer utils.Disconnect(account)

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("read error: ", err)
			}
			return
		}

		switch messageType {
		case websocket.TextMessage:
			err := service.ProcessMessage(account, message)
			if err != nil {
				logger.Error("process message error: ", err)
				c.WriteJSON(fiber.Map{"message": err})
			}
		case websocket.PingMessage:
			c.WriteMessage(websocket.PongMessage, []byte{})
		case websocket.PongMessage:
		default:
			logger.Info("websocket message received of type: ", messageType)
		}
	}
}
