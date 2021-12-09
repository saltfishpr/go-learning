// @description: 注册登录注销
// @file: auth.go
// @date: 2021/11/21

package v1

import (
	"time"

	"learning/config"
	"learning/internal/constant/e"
	"learning/internal/model"
	"learning/internal/service"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Register is a function to sign up
// @Summary Create an account
// @Description Create an account
// @Tags auth
// @Accept json
// @Produce json
// @Success 201
// @Failure 400 {object} ResponseHTTP{}
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	if user.Nickname == nil {
		user.Nickname = user.Account
	}
	err := service.CreateUser(user)
	if err != nil {
		logger.Error("create user error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.ExistAccount))
	}

	return c.SendStatus(fiber.StatusCreated)
}

func Login(c *fiber.Ctx) error {
	account := c.FormValue("account")
	password := c.FormValue("password")
	if len(account) == 0 || len(password) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}

	user, err := service.GetUserByAccount(account)
	if err != nil || *(user.Password) != password {
		logger.Error("login error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.LoginFailed))
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
