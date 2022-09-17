// @description: 处理 users 接口相关的 http 请求
// @file: user.go
// @date: 2021/11/21

package user

import (
	"learning/internal/constant/e"
	"learning/internal/log"
	"learning/internal/model"
	"learning/internal/service"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service.IUser
	logger  *log.Logger
}

func New(service service.IUser, logger *log.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	if len(username) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing username")))
	}

	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		h.logger.Error("get user error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get user failed")))
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		h.logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	username := c.Params("username")
	if len(username) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing username")))
	}

	err := h.service.UpdateUser(user)
	if err != nil {
		h.logger.Error("update user error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("update user failed")))
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	username := c.Params("username")
	if len(username) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing username")))
	}

	err := h.service.DeleteUserByUsername(username)
	if err != nil {
		h.logger.Error("delete user error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("delete user failed")))
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.logger.Error("get all users error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get all users failed")))
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *Handler) Join(c *fiber.Ctx) error {
	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	username := utils.MustGetUsernameFromCtx(c)

	err := h.service.JoinHub(username, hid)
	if err != nil {
		h.logger.Errorf("user %s join hub %s error: %s", username, hid, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("join hub failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) Leave(c *fiber.Ctx) error {
	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	username := utils.MustGetUsernameFromCtx(c)

	err := h.service.LeaveHub(username, hid)
	if err != nil {
		h.logger.Errorf("user %s leave hub %s error: %s", username, hid, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("leave hub failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) Joined(c *fiber.Ctx) error {
	username := utils.MustGetUsernameFromCtx(c)

	hubs, err := h.service.GetJoinedHubs(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			e.Failed(
				e.Error, e.WithMessage("get joined hubs failed"),
			),
		)
	}
	return c.Status(fiber.StatusOK).JSON(hubs)
}

func (h *Handler) Follow(c *fiber.Ctx) error {
	username := c.Query("username")
	if len(username) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing username")))
	}

	curUsername := utils.MustGetUsernameFromCtx(c)

	err := h.service.FollowUser(curUsername, username)
	if err != nil {
		h.logger.Errorf("user %s follow friend %s error: %s", curUsername, username, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("follow user failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) Unfollow(c *fiber.Ctx) error {
	username := c.Query("username")
	if len(username) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing username")))
	}

	curUsername := utils.MustGetUsernameFromCtx(c)

	err := h.service.UnfollowUser(curUsername, username)
	if err != nil {
		h.logger.Errorf("user %s unfollow friend %s error: %s", curUsername, username, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("unfollow user failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) Following(c *fiber.Ctx) error {
	username := utils.MustGetUsernameFromCtx(c)

	friends, err := h.service.GetFollowingUsers(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			e.Failed(
				e.Error, e.WithMessage("get following users failed"),
			),
		)
	}
	return c.Status(fiber.StatusOK).JSON(friends)
}
