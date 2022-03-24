// @description: 注册登录注销
// @file: auth.go
// @date: 2021/11/21

package v1

import (
	"errors"

	"learning/internal/common/rediscache"
	"learning/internal/constant"
	"learning/internal/constant/e"
	"learning/internal/model"
	"learning/internal/service"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = utils.NewValidate()

func Login(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	userX := new(model.User)
	if err := c.BodyParser(userX); err != nil {
		logger.Error("parse body error: ", err)
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

	userService := service.NewUser(conn)
	user, err := userService.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := userService.CreateUser(userX); err != nil {
				logger.Error("create user error: ", err)
				return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.ExistUsername))
			}
		} else {
			logger.Error("get user error: ", err)
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

func Refresh(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)

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
		logger.Error("delete refresh token error: ", err)
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
