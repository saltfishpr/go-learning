// @description: 处理 hubs 接口相关的 http 请求
// @file: hubs.go
// @date: 2021/11/21

package v1

import (
	"net/http"

	"learning/internal/constant/e"
	"learning/internal/model"
	"learning/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func CreateHub(c *fiber.Ctx) error {
	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		zap.S().Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}

	err := service.CreateHub(hub)
	if err != nil {
		zap.S().Error("create hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.ExistHub))
	}

	return c.SendStatus(fiber.StatusCreated)
}

func GetAllHubs(c *fiber.Ctx) error {
	hubs, err := service.GetAllHubs()
	if err != nil {
		zap.S().Error("get hubs error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("get all hubs failed")))
	}

	return c.Status(fiber.StatusOK).JSON(hubs)
}

func UpdateHub(c *fiber.Ctx) error {
	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		zap.S().Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	if hub.HID == nil {
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	err := service.UpdateHub(hub)
	if err != nil {
		zap.S().Error("update hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("update hub failed")))
	}

	return c.SendStatus(http.StatusOK)
}

func DeleteHub(c *fiber.Ctx) error {
	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	err := service.DeleteUserByAccount(hid)
	if err != nil {
		zap.S().Error("delete hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("delete hub failed")))
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetHubInfo(c *fiber.Ctx) error {
	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	hub, err := service.GetHubByHID(hid)
	if err != nil {
		zap.S().Error("get hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("get hub failed")))
	}

	return c.Status(fiber.StatusOK).JSON(hub)
}
