// @description: 用户服务
// @file: user.go
// @date: 2021/11/22

package service

import (
	"learning/internal/data"
	"learning/internal/model"

	"github.com/jinzhu/copier"
)

type IUser interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUserByUsername(username string) error
	GetAllUsers() ([]*model.User, error)

	JoinHub(username string, hid string) error
	LeaveHub(username string, hid string) error
	GetJoinedHubs(username string) ([]*model.Hub, error)

	FollowUser(username string, targetUsername string) error
	UnfollowUser(username string, targetUsername string) error
	GetFollowingUsers(username string) ([]*model.User, error)
}

type User struct {
	connection data.Connection
}

var _ IUser = (*User)(nil)

func NewUser(connection data.Connection) *User {
	return &User{
		connection: connection,
	}
}

func (u *User) CreateUser(user *model.User) error {
	userEntity := new(data.User)
	copier.Copy(userEntity, user)

	return u.connection.CreateUser(userEntity)
}

func (u *User) GetUserByUsername(username string) (*model.User, error) {
	userEntity, err := u.connection.GetUserByUsername(username)
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

	return u.connection.UpdateUser(userEntity)
}

func (u *User) DeleteUserByUsername(username string) error {
	return u.connection.DeleteUserByUsername(username)
}

func (u *User) GetAllUsers() ([]*model.User, error) {
	userEntities, err := u.connection.GetAllUsers()
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, len(userEntities))
	copier.Copy(&users, &userEntities)

	return users, nil
}

func (u *User) JoinHub(username string, hid string) error {
	userEntity, err := u.connection.GetUserByUsername(username)
	if err != nil {
		return err
	}

	hubEntity, err := u.connection.GetHubByHID(hid)
	if err != nil {
		return err
	}

	return u.connection.UserJoinHub(userEntity, hubEntity)
}

func (u *User) LeaveHub(username string, hid string) error {
	userEntity, err := u.connection.GetUserByUsername(username)
	if err != nil {
		return err
	}

	hubEntity, err := u.connection.GetHubByHID(hid)
	if err != nil {
		return err
	}

	return u.connection.UserLeaveHub(userEntity, hubEntity)
}

func (u *User) GetJoinedHubs(username string) ([]*model.Hub, error) {
	userEntity, err := u.connection.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	hubEntities, err := u.connection.GetUserJoinedHubs(userEntity)
	if err != nil {
		return nil, err
	}
	hubs := make([]*model.Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)

	return hubs, nil
}

func (u *User) FollowUser(username string, targetUsername string) error {
	userEntity, err := u.connection.GetUserByUsername(username)
	if err != nil {
		return err
	}

	friendEntity, err := u.connection.GetUserByUsername(targetUsername)
	if err != nil {
		return err
	}

	return u.connection.UserFollowUser(userEntity, friendEntity)
}

func (u *User) UnfollowUser(username string, targetUsername string) error {
	userEntity, err := u.connection.GetUserByUsername(username)
	if err != nil {
		return err
	}

	friendEntity, err := u.connection.GetUserByUsername(targetUsername)
	if err != nil {
		return err
	}

	return u.connection.UserUnfollowUser(userEntity, friendEntity)
}

func (u *User) GetFollowingUsers(username string) ([]*model.User, error) {
	userEntity, err := u.connection.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	friendEntities, err := u.connection.GetUserFollowingUsers(userEntity)
	if err != nil {
		return nil, err
	}
	friends := make([]*model.User, len(friendEntities))
	copier.Copy(&friends, &friendEntities)

	return friends, nil
}
