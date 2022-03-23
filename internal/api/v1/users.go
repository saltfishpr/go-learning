// @description: 处理 users 接口相关的 http 请求
// @file: users.go
// @date: 2021/11/21

package v1

import (
	"learning/internal/constant/e"
	"learning/internal/model"
	"learning/internal/service"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	userService := service.NewUser(conn)
	users, err := userService.GetUserList()
	if err != nil {
		logger.Error("get all users error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get all users failed")))
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func UpdateUser(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("parse body error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidParams))
	}
	if user.Account == nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing account")))
	}
	userService := service.NewUser(conn)
	err := userService.UpdateUser(user)
	if err != nil {
		logger.Error("update user error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("update user failed")))
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteUser(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	account := c.Query("account")
	if len(account) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing account")))
	}
	userService := service.NewUser(conn)
	err := userService.DeleteUserByAccount(account)
	if err != nil {
		logger.Error("delete user error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("delete user failed")))
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUserInfo(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	account := c.Params("account")
	if len(account) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing account")))
	}

	userService := service.NewUser(conn)
	user, err := userService.GetUserByAccount(account)
	if err != nil {
		logger.Error("get user error: ", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("get user failed")))
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func Join(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	account := utils.MustGetUserAccountFromCtx(c)

	userService := service.NewUser(conn)
	err := userService.JoinHub(account, hid)
	if err != nil {
		logger.Errorf("user %s join hub %s error: %s", account, hid, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("join hub failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}

func Joined(c *fiber.Ctx) error {
	conn := utils.MustGetConnectionFromContext(c)

	account := utils.MustGetUserAccountFromCtx(c)

	userService := service.NewUser(conn)
	hubs, err := userService.GetJoinedHubs(account)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			e.Failed(
				e.Error, e.WithMessage("get joined hubs failed"),
			),
		)
	}
	return c.Status(fiber.StatusOK).JSON(hubs)
}

func Leave(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	hid := c.Query("hid")
	if len(hid) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing hid")))
	}

	account := utils.MustGetUserAccountFromCtx(c)

	userService := service.NewUser(conn)
	err := userService.LeaveHub(account, hid)
	if err != nil {
		logger.Errorf("user %s leave hub %s error: %s", account, hid, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("leave hub failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}

func Follow(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	friendAccount := c.Query("account")
	if len(friendAccount) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing account")))
	}

	account := utils.MustGetUserAccountFromCtx(c)

	userService := service.NewUser(conn)
	err := userService.FollowUser(account, friendAccount)
	if err != nil {
		logger.Errorf("user %s follow friend %s error: %s", account, friendAccount, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("follow user failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}

func Following(c *fiber.Ctx) error {
	conn := utils.MustGetConnectionFromContext(c)

	account := utils.MustGetUserAccountFromCtx(c)

	userService := service.NewUser(conn)
	friends, err := userService.GetFollowingUsers(account)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			e.Failed(
				e.Error, e.WithMessage("get following users failed"),
			),
		)
	}
	return c.Status(fiber.StatusOK).JSON(friends)
}

func Unfollow(c *fiber.Ctx) error {
	logger := utils.MustGetLoggerFromContext(c)
	conn := utils.MustGetConnectionFromContext(c)

	friendAccount := c.Query("account")
	if len(friendAccount) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(e.Failed(e.InvalidParams, e.WithMessage("missing account")))
	}

	account := utils.MustGetUserAccountFromCtx(c)

	userService := service.NewUser(conn)
	err := userService.UnfollowUser(account, friendAccount)
	if err != nil {
		logger.Errorf("user %s unfollow friend %s error: %s", account, friendAccount, err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(e.Failed(e.Error, e.WithMessage("unfollow user failed")))
	}
	return c.SendStatus(fiber.StatusOK)
}
