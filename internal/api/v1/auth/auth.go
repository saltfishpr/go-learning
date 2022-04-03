// @description: 注册登录注销
// @file: auth.go
// @date: 2021/11/21

package auth

import (
	"errors"

	"learning/internal/common/rediscache"
	"learning/internal/constant"
	"learning/internal/constant/e"
	"learning/internal/log"
	"learning/internal/model"
	"learning/internal/service"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = utils.NewValidate()

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

func (h *Handler) Login(c *fiber.Ctx) error {
	userX := new(model.User)
	if err := c.BodyParser(userX); err != nil {
		h.logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	err := validate.Struct(userX)
	if err != nil {
		utils.LogValidateErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	username := *userX.Username
	password := *userX.Password
	if userX.Nickname == nil {
		userX.Nickname = userX.Username
	}

	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := h.service.CreateUser(userX); err != nil {
				h.logger.Error("create user error: ", err)
				return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.ExistUsername))
			}
		} else {
			h.logger.Error("get user error: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error))
		}
	} else if *(user.Password) != password {
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.LoginFailed))
	}

	t, err := utils.GenerateTokenPair(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("Generate token failed.")))
	}
	return c.JSON(t)
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	claims, err := utils.VerifyToken(c.Query("refresh_token"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized))
	}
	var username string
	if err := rediscache.Get(constant.RefreshTokenPrefix+claims["jti"].(string), &username); err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(e.Failed(e.Unauthorized, e.WithMessage("Invalid refresh token.")))
	}
	if err := rediscache.Del(constant.RefreshTokenPrefix + claims["jti"].(string)); err != nil {
		h.logger.Error("delete refresh token error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("Generate token failed.")))
	}
	t, err := utils.GenerateTokenPair(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("Generate token failed.")))
	}
	return c.JSON(t)
}
