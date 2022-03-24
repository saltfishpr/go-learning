// @description: 处理 hubs 接口相关的 http 请求
// @file: hubs.go
// @date: 2021/11/21

package v1

import (
	"net/http"

	"learning/internal/constant/e"
	"learning/internal/model"
	"learning/internal/service"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateHub(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}

	hubService := service.NewHub(conn)
	if err := hubService.CreateHub(hub); err != nil {
		logger.Error("create hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.ExistHub))
	}

	return c.SendStatus(fiber.StatusCreated)
}

func GetHub(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	hubService := service.NewHub(conn)
	hub, err := hubService.GetHubByHID(hid)
	if err != nil {
		logger.Error("get hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get hub failed")))
	}

	return c.Status(fiber.StatusOK).JSON(hub)
}

func UpdateHub(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	hubService := service.NewHub(conn)
	err := hubService.UpdateHub(hub)
	if err != nil {
		logger.Error("update hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("update hub failed")))
	}

	return c.SendStatus(http.StatusOK)
}

func DeleteHub(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	hubService := service.NewHub(conn)
	err := hubService.DeleteHubByHID(hid)
	if err != nil {
		logger.Error("delete hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("delete hub failed")))
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetHubs(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	hubService := service.NewHub(conn)
	hubs, err := hubService.GetAllHubs()
	if err != nil {
		logger.Error("get hubs error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get all hubs failed")))
	}

	return c.Status(fiber.StatusOK).JSON(hubs)
}
