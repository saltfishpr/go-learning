// @description: 用户服务
// @file: user.go
// @date: 2021/11/22

package service

import (
	"learning/internal/data"
	"learning/internal/model"

	"github.com/jinzhu/copier"
)

type User struct {
	conn data.Connection
}

func NewUser(conn data.Connection) *User {
	return &User{conn: conn}
}

func (u *User) CreateUser(user *model.User) error {
	userEntity := new(data.User)
	copier.Copy(userEntity, user)
	return u.conn.CreateUser(userEntity)
}

func (u *User) GetUserByAccount(account string) (*model.User, error) {
	userEntity, err := u.conn.GetUserByAccount(account)
	if err != nil {
		return nil, err
	}
	user := new(model.User)
	copier.Copy(user, userEntity)
	return user, nil
}

func (u *User) UpdateUser(user *model.User) error {
	userEntity := new(data.User)
	copier.Copy(userEntity, user)
	return u.conn.UpdateUser(userEntity)
}

func (u *User) DeleteUserByAccount(account string) error {
	return u.conn.DeleteUserByAccount(account)
}

func (u *User) GetUserList() ([]*model.User, error) {
	userEntities, err := u.conn.GetAllUsers()
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, len(userEntities))
	copier.Copy(&users, &userEntities)

	return users, nil
}

func (u *User) JoinHub(account string, hid string) error {
	userEntity, err := u.conn.GetUserByAccount(account)
	if err != nil {
		return err
	}

	hubEntity, err := u.conn.GetHubByHID(hid)
	if err != nil {
		return err
	}

	return u.conn.UserJoinHub(userEntity, hubEntity)
}

func (u *User) LeaveHub(account string, hid string) error {
	userEntity, err := u.conn.GetUserByAccount(account)
	if err != nil {
		return err
	}

	hubEntity, err := u.conn.GetHubByHID(hid)
	if err != nil {
		return err
	}

	return u.conn.UserLeaveHub(userEntity, hubEntity)
}

func (u *User) GetJoinedHubs(account string) ([]*model.Hub, error) {
	userEntity, err := u.conn.GetUserByAccount(account)
	if err != nil {
		return nil, err
	}

	hubEntities, err := u.conn.GetUserJoinedHubs(userEntity)
	if err != nil {
		return nil, err
	}
	hubs := make([]*model.Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)

	return hubs, nil
}

func (u *User) FollowUser(account string, friendAccount string) error {
	userEntity, err := u.conn.GetUserByAccount(account)
	if err != nil {
		return err
	}

	friendEntity, err := u.conn.GetUserByAccount(friendAccount)
	if err != nil {
		return err
	}

	return u.conn.UserFollowUser(userEntity, friendEntity)
}

func (u *User) UnfollowUser(account string, friendAccount string) error {
	userEntity, err := u.conn.GetUserByAccount(account)
	if err != nil {
		return err
	}

	friendEntity, err := u.conn.GetUserByAccount(friendAccount)
	if err != nil {
		return err
	}

	return u.conn.UserUnfollowUser(userEntity, friendEntity)
}

func (u *User) GetFollowingUsers(account string) ([]*model.User, error) {
	userEntity, err := u.conn.GetUserByAccount(account)
	if err != nil {
		return nil, err
	}

	friendEntities, err := u.conn.GetUserFollowingUsers(userEntity)
	if err != nil {
		return nil, err
	}
	friends := make([]*model.User, len(friendEntities))
	copier.Copy(&friends, &friendEntities)

	return friends, nil
}
