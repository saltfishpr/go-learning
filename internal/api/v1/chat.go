// @description: 处理用户发送的消息
// @file: chat.go
// @date: 2021/11/22

package v1

import (
	"context"

	"learning/internal/common/connstorage"
	"learning/internal/common/gocache"
	"learning/internal/service"
	"learning/internal/utils"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ChatAuth(c *fiber.Ctx) error {
	sid := utils.GenerateSID().String()
	account := utils.MustGetUserAccountFromCtx(c)
	gocache.Set(sid, account, gocache.NoExpiration)
	return c.Status(fiber.StatusUpgradeRequired).JSON(fiber.Map{"sid": sid})
}

func ChatHandler(c *websocket.Conn) {
	sid := c.Locals("sid").(string)
	connstorage.Set(sid, c)
	defer connstorage.Del(sid)

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
			ctx := context.WithValue(context.Background(), "account", c.Locals("account"))
			err := service.ProcessMessage(ctx, message)
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
