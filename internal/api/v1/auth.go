// @description: 注册登录注销
// @file: auth.go
// @date: 2021/11/21

package v1

import (
	"errors"
	"time"

	"learning/config"
	"learning/internal/constant/e"
	"learning/internal/model"
	"learning/internal/service"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// Login is a function to sign or sign up
// @Summary Sign in or sign up.
// @Description Sign in if account exists. Otherwise, sign up. If success, return a jwt token.
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map{"token": t, "expire_at": expireAt}
// @Success 201
// @Failure 400 {object} e.ErrorResult
// @Router /register [post]
func Login(c *fiber.Ctx) error {
	userX := new(model.User)
	if err := c.BodyParser(userX); err != nil {
		logger.Error("parse body error: ", err)
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

	expireAt := time.Now().Add(config.TokenExpireTime).Unix()
	claims := jwt.MapClaims{
		"account":   account,
		"expire_at": expireAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		logger.Error("sign token error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error, e.WithMessage("generate token failed")))
	}

	return c.JSON(fiber.Map{"token": t, "expire_at": expireAt})
}

func Check(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
