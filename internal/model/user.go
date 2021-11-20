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

	Hubs    []*Hub  `json:"hubs,omitempty"`
	Friends []*User `json:"friends,omitempty"`
}

func CreateUser(user *User) error {
	userEntity := new(data.User)
	copier.Copy(userEntity, user)
	return data.CreateUser(userEntity)
}

func ReadAllUsers() ([]*User, error) {
	userEntities, err := data.ReadAllUsers()
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

func ReadUserByAccount(account string) (*User, error) {
	userEntity, err := data.ReadUserByAccount(account)
	if err != nil {
		return nil, err
	}
	user := new(User)
	copier.Copy(user, userEntity)

	hubEntities, err := data.GetJoinedHubs(userEntity)
	if err != nil {
		return nil, err
	}
	hubs := make([]*Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)
	user.Hubs = hubs
	return user, nil
}
