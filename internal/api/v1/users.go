// @description: 处理用户的增删改查操作
// @file: users.go
// @date: 2021/11/21

package v1

import (
	"github.com/gofiber/fiber/v2"

	"learning/internal/model"
	"learning/logger"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := model.ReadAllUsers()
	if err != nil {
		logger.Error("read users error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "数据错误"})
	}
	if user.Account == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少用户ID"})
	}

	err := model.UpdateUser(user)
	if err != nil {
		logger.Error("update user error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteUser(c *fiber.Ctx) error {
	account := c.Query("account")
	if len(account) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少用户ID"})
	}

	err := model.DeleteUserByAccount(account)
	if err != nil {
		logger.Error("delete user error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUserInfo(c *fiber.Ctx) error {
	account := c.Params("account")
	if len(account) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少用户ID"})
	}

	user, err := model.ReadUserByAccount(account)
	if err != nil {
		logger.Error("read user error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "没有此用户"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
