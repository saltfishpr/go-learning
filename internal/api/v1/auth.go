// @description: 注册登录注销
// @file: auth.go
// @date: 2021/11/21

package v1

import (
	"errors"

	"learning/config"
	"learning/internal/common/rediscache"
	"learning/internal/constant/e"
	"learning/internal/model"
	"learning/internal/service"
	"learning/internal/utils"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = utils.NewValidate()

// Login is a function to sign or sign up
// @Summary Sign in or sign up.
// @Description Sign in if account exists. Otherwise, sign up. If success, return a jwt token.
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map{"token": token}
// @Success 201
// @Failure 400 {object} e.ErrorResult
// @Router /register [post]
func Login(c *fiber.Ctx) error {
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
	account := *userX.Account
	password := *userX.Password
	if userX.Nickname == nil {
		userX.Nickname = userX.Account
	}

	user, err := service.GetUserByAccount(account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := service.CreateUser(userX); err != nil {
				logger.Error("create user error: ", err)
				return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.ExistAccount))
			}
		} else {
			logger.Error("get user error: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error))
		}
	} else {
		if *(user.Password) != password {
			logger.Error("login error: ", err)
			return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.LoginFailed))
		}
	}

	t, err := utils.GenerateToken(account)
	rt, err := utils.GenerateRefreshToken(account)
	if err != nil {
		logger.Error("sign token error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("generate token failed")))
	}

	return c.JSON(fiber.Map{"token": t, "refresh_token": rt})
}

func Refresh(c *fiber.Ctx) error {
	claims, err := utils.VerifyToken(c.Query("refresh_token"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized))
	}
	var account string
	if err := rediscache.Get(config.RefreshTokenPrefix+claims["jti"].(string), &account); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized, e.WithMessage("Refresh Token Invalid")))
	}
	if err := rediscache.Del(config.RefreshTokenPrefix + claims["jti"].(string)); err != nil {
		logger.Error("delete refresh token error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("generate token failed")))
	}
	t, err := utils.GenerateToken(account)
	rt, err := utils.GenerateRefreshToken(account)
	if err != nil {
		logger.Error("sign token error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("generate token failed")))
	}
	return c.JSON(fiber.Map{"token": t, "refresh_token": rt})
}

func Check(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
