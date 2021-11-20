// @description: 用户数据结构与持久化
// @file: user.go
// @date: 2021/11/18

package model

import (
	"github.com/jinzhu/copier"

	"learning/internal/data"
)

type User struct {
	Account  *string `json:"account" validate:"required"`
	Password *string `json:"password" validate:"required"`
	Nickname *string `json:"nickname" validate:"required"`
	Address  *string `json:"address,omitempty"`
}

func CreateUser(user *User) error {
	userEntity := new(data.User)
	copier.Copy(userEntity, user)
	return data.CreateUser(userEntity)
}

func GetAllUsers() ([]*User, error) {
	userEntities, err := data.GetAllUsers()
	if err != nil {
		return nil, err
	}
	users := make([]*User, len(userEntities))
	copier.Copy(&users, &userEntities)

	return users, nil
}

func UpdateUser(user *User) error {
	userEntity := new(data.User)
	copier.Copy(userEntity, user)
	return data.UpdateUser(userEntity)
}

func DeleteUserByAccount(account string) error {
	return data.DeleteUserByAccount(account)
}

func GetUserByAccount(account string) (*User, error) {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return nil, err
	}
	user := new(User)
	copier.Copy(user, userEntity)

	return user, nil
}

func JoinHub(account string, hid string) error {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return err
	}

	hubEntity, err := data.GetHubByHID(hid)
	if err != nil {
		return err
	}

	return data.JoinHub(userEntity, hubEntity)
}

func GetJoinedHubs(account string) ([]*Hub, error) {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return nil, err
	}

	hubEntities, err := data.GetJoinedHubs(userEntity)
	if err != nil {
		return nil, err
	}
	hubs := make([]*Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)

	return hubs, nil
}

func LeaveHub(account string, hid string) error {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return err
	}

	hubEntity, err := data.GetHubByHID(hid)
	if err != nil {
		return err
	}

	return data.LeaveHub(userEntity, hubEntity)
}

func FollowUser(account string, friendAccount string) error {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return err
	}

	friendEntity, err := data.GetUserByAccount(friendAccount)
	if err != nil {
		return err
	}

	return data.FollowUser(userEntity, friendEntity)
}

func GetFollowingUsers(account string) ([]*User, error) {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return nil, err
	}

	friendEntities, err := data.GetFollowingUsers(userEntity)
	if err != nil {
		return nil, err
	}
	friends := make([]*User, len(friendEntities))
	copier.Copy(&friends, &friendEntities)

	return friends, nil
}

func UnfollowUser(account string, friendAccount string) error {
	userEntity, err := data.GetUserByAccount(account)
	if err != nil {
		return err
	}

	friendEntity, err := data.GetUserByAccount(friendAccount)
	if err != nil {
		return err
	}

	return data.UnfollowUser(userEntity, friendEntity)
}
