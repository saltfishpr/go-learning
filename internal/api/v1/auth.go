// @description: 注册登录注销
// @file: auth.go
// @date: 2021/11/21

package v1

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"learning/config"
	"learning/internal/model"
	"learning/internal/service"
	"learning/logger"
)

func Register(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "数据错误"})
	}

	err := service.CreateUser(user)
	if err != nil {
		logger.Error("create user error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func Login(c *fiber.Ctx) error {
	account := c.FormValue("account")
	password := c.FormValue("password")

	user, err := service.GetUserByAccount(account)
	if err != nil || *(user.Password) != password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "用户名或密码错误"})
	}

	expireAt := time.Now().Add(config.TokenExpireTime).Unix()
	claims := jwt.MapClaims{
		"account":   account,
		"expire_at": expireAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t, "expire_at": expireAt})
}

func Logout(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
