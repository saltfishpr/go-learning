// @description: 处理 hubs 接口相关的 http 请求
// @file: hubs.go
// @date: 2021/11/21

package v1

import (
	"net/http"

	"learning/internal/model"
	"learning/internal/service"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
)

func CreateHub(c *fiber.Ctx) error {
	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "数据错误"})
	}

	err := service.CreateHub(hub)
	if err != nil {
		logger.Error("create hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func GetAllHubs(c *fiber.Ctx) error {
	hubs, err := service.GetAllHubs()
	if err != nil {
		logger.Error("get hubs error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.Status(fiber.StatusOK).JSON(hubs)
}

func UpdateHub(c *fiber.Ctx) error {
	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "数据错误"})
	}
	if hub.HID == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少聊天室ID"})
	}

	err := service.UpdateHub(hub)
	if err != nil {
		logger.Error("update hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.SendStatus(http.StatusOK)
}

func DeleteHub(c *fiber.Ctx) error {
	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少聊天室ID"})
	}

	err := service.DeleteUserByAccount(hid)
	if err != nil {
		logger.Error("delete hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetHubInfo(c *fiber.Ctx) error {
	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少聊天室ID"})
	}

	hub, err := service.GetHubByHID(hid)
	if err != nil {
		logger.Error("get hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.Status(fiber.StatusOK).JSON(hub)
}
