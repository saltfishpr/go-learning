// @file: users.go
// @date: 2021/11/18

// Package users 处理用户的增删改查操作
package users

import (
	"github.com/gofiber/fiber/v2"

	"learning/internal/model"
	"learning/logger"
)

func Create(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("数据错误")
	}

	err := model.CreateUser(user)
	if err != nil {
		logger.Error("create user error: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("服务器出现错误")
	}

	return c.SendStatus(fiber.StatusOK)
}

func Read(c *fiber.Ctx) error {
	account := c.Query("account")
	if len(account) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("缺少用户ID")
	}

	user, err := model.ReadUserByAccount(account)
	if err != nil {
		logger.Error("read user error: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("服务器出现错误")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func ReadAll(c *fiber.Ctx) error {
	users, err := model.ReadAllUsers()
	if err != nil {
		logger.Error("read user error: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("服务器出现错误")
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func Update(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("数据错误")
	}
	if user.Account == nil {
		logger.Debugf("%v", user)
		return c.Status(fiber.StatusBadRequest).SendString("缺少用户ID")
	}

	err := model.UpdateUser(user)
	if err != nil {
		logger.Error("update user error: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("服务器出现错误")
	}

	return c.SendStatus(fiber.StatusOK)
}

func Delete(c *fiber.Ctx) error {
	account := c.Query("account")
	if len(account) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("缺少用户ID")
	}

	err := model.DeleteUserByAccount(account)
	if err != nil {
		logger.Error("delete user error: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("服务器出现错误")
	}
	return c.SendStatus(fiber.StatusOK)
}
