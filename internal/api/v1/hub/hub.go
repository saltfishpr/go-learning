// @description: 处理 hubs 接口相关的 http 请求
// @file: hub.go
// @date: 2021/11/21

package hub

import (
	"net/http"

	"learning/internal/constant/e"
	"learning/internal/log"
	"learning/internal/model"
	"learning/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service.IHub
	logger  *log.Logger
}

func New(service service.IHub, logger *log.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateHub(c *fiber.Ctx) error {
	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		h.logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}

	if err := h.service.CreateHub(hub); err != nil {
		h.logger.Error("create hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.ExistHub))
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) GetHub(c *fiber.Ctx) error {
	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	hub, err := h.service.GetHubByHID(hid)
	if err != nil {
		h.logger.Error("get hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get hub failed")))
	}

	return c.Status(fiber.StatusOK).JSON(hub)
}

func (h *Handler) UpdateHub(c *fiber.Ctx) error {
	hub := new(model.Hub)
	if err := c.BodyParser(hub); err != nil {
		h.logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	err := h.service.UpdateHub(hub)
	if err != nil {
		h.logger.Error("update hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("update hub failed")))
	}

	return c.SendStatus(http.StatusOK)
}

func (h *Handler) DeleteHub(c *fiber.Ctx) error {
	hid := c.Params("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	err := h.service.DeleteHubByHID(hid)
	if err != nil {
		h.logger.Error("delete hub error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("delete hub failed")))
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) GetHubs(c *fiber.Ctx) error {
	hubs, err := h.service.GetAllHubs()
	if err != nil {
		h.logger.Error("get hubs error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get all hubs failed")))
	}

	return c.Status(fiber.StatusOK).JSON(hubs)
}
