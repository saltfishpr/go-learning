// @description: 处理用户发送的消息
// @file: chat.go
// @date: 2021/11/22

package chat

import (
	"context"

	"learning/internal/common/connstorage"
	"learning/internal/constant/e"
	"learning/internal/log"
	"learning/internal/model"
	"learning/internal/service"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type Handler struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) ChatAuth(c *fiber.Ctx) error {
	username := utils.MustGetUsernameFromCtx(c)
	token, err := utils.GenerateDisposableToken(username)
	if err != nil {
		zap.S().Error("sign token error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("generate token failed")))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func (h *Handler) ChatHandler(c *websocket.Conn) {
	username := c.Locals("username").(string)
	connstorage.Set(username, c)
	defer connstorage.Del(username)

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				zap.S().Error("read error: ", err)
			}
			return
		}

		switch messageType {
		case websocket.TextMessage:
			err := service.ProcessMessage(context.TODO(), message)
			if err != nil {
				zap.S().Error("process message error: ", err)
				c.WriteJSON(fiber.Map{"message": err})
			}
		case websocket.PingMessage:
			c.WriteMessage(websocket.PongMessage, []byte{})
		case websocket.PongMessage:
		default:
			zap.S().Info("websocket message received of type: ", messageType)
		}
	}
}

func (h *Handler) GetMessages(c *fiber.Ctx) error {
	topic := c.Params("topic")
	offset := c.Query("offset")
	limit := c.Query("limit")
	if len(topic) == 0 || len(offset) == 0 || len(limit) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	query := &model.MessagesPaginationRequest{
		Topic:  topic,
		Offset: cast.ToInt(offset),
		Limit:  cast.ToInt(limit),
	}
	messages, err := service.GetMessagesPagination(utils.MustGetUsernameFromCtx(c), query)
	if err != nil {
		zap.S().Error("get messages error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error))
		// TODO: Add ErrorCode
	}
	return c.Status(fiber.StatusOK).JSON(messages)
}
