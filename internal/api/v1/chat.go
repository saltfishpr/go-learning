// @description: 处理用户发送的消息
// @file: chat.go
// @date: 2021/11/22

package v1

import (
	"context"

	"learning/internal/common/connstorage"
	"learning/internal/constant/e"
	"learning/internal/service"
	"learning/internal/utils"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ChatAuth(c *fiber.Ctx) error {
	// TODO: generate token with jti
	account := utils.MustGetUserAccountFromCtx(c)
	token, err := utils.GenerateDisposableToken(account)
	if err != nil {
		logger.Error("sign token error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("generate token failed")))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func ChatHandler(c *websocket.Conn) {
	account := c.Locals("account").(string)
	connstorage.Set(account, c)
	defer connstorage.Del(account)

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
			ctx := context.WithValue(context.Background(), "account", account)
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
