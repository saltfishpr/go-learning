// @description: 处理用户的增删改查操作
// @file: users.go
// @date: 2021/11/21

package v1

import (
	"github.com/gofiber/fiber/v2"

	"learning/internal/model"
	"learning/internal/utils"
	"learning/logger"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := model.GetAllUsers()
	if err != nil {
		logger.Error("get users error: ", err)
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

	user, err := model.GetUserByAccount(account)
	if err != nil {
		logger.Error("get user error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "没有此用户"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func Join(c *fiber.Ctx) error {
	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少聊天室ID"})
	}

	account := utils.GetUserAccountFromCtx(c)
	err := model.JoinHub(account, hid)
	if err != nil {
		logger.Errorf("user %s join hub %s error: %s", account, hid, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}
	return c.SendStatus(fiber.StatusOK)
}

func Joined(c *fiber.Ctx) error {
	account := utils.GetUserAccountFromCtx(c)
	hubs, err := model.GetJoinedHubs(account)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}
	return c.Status(fiber.StatusOK).JSON(hubs)
}

func Leave(c *fiber.Ctx) error {
	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少聊天室ID"})
	}

	account := utils.GetUserAccountFromCtx(c)
	err := model.LeaveHub(account, hid)
	if err != nil {
		logger.Errorf("user %s leave hub %s error: %s", account, hid, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}
	return c.SendStatus(fiber.StatusOK)
}

func Follow(c *fiber.Ctx) error {
	friendAccount := c.Query("account")
	if len(friendAccount) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少用户ID"})
	}

	account := utils.GetUserAccountFromCtx(c)
	err := model.FollowUser(account, friendAccount)
	if err != nil {
		logger.Errorf("user %s follow friend %s error: %s", account, friendAccount, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}
	return c.SendStatus(fiber.StatusOK)
}

func Following(c *fiber.Ctx) error {
	account := utils.GetUserAccountFromCtx(c)
	friends, err := model.GetFollowingUsers(account)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}
	return c.Status(fiber.StatusOK).JSON(friends)
}

func Unfollow(c *fiber.Ctx) error {
	friendAccount := c.Query("account")
	if len(friendAccount) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少用户ID"})
	}

	account := utils.GetUserAccountFromCtx(c)
	err := model.UnfollowUser(account, friendAccount)
	if err != nil {
		logger.Errorf("user %s unfollow friend %s error: %s", account, friendAccount, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "服务器出现错误"})
	}
	return c.SendStatus(fiber.StatusOK)
}
