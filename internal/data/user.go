// @description: 用户表字段
// @file: user.go
// @date: 2021/11/19

package data

import (
	"gorm.io/plugin/soft_delete"
)

type User struct {
	Base
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:user_udx_delete"`

	Account  *string `gorm:"size:32;not null;uniqueIndex:user_udx_delete"`
	Password *string `gorm:"size:32;not null"`
	Phone    *string `gorm:"size:32"`
	Email    *string `gorm:"size:32"`
	Nickname *string `gorm:"size:32;not null"`
	Address  *string `gorm:"size:128"`

	Hubs    []*Hub  `gorm:"many2many:user_hub"`
	Friends []*User `gorm:"many2many:user_friends"`
}

func CreateUser(user *User) error {
	res := db.Create(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetAllUsers() ([]*User, error) {
	var users []*User

	res := db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func UpdateUser(user *User) error {
	res := db.Model(user).Where("account = ?", user.Account).Updates(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteUserByAccount(account string) error {
	res := db.Where("account = ?", account).Delete(&User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetUserByAccount(account string) (*User, error) {
	user := new(User)

	res := db.Where("account = ?", account).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func JoinHub(user *User, hub *Hub) error {
	return db.Model(user).Association("Hubs").Append(hub)
}

func GetJoinedHubs(user *User) ([]*Hub, error) {
	var hubs []*Hub
	err := db.Model(&user).Association("Hubs").Find(&hubs)
	if err != nil {
		return nil, err
	}
	return hubs, nil
}

func LeaveHub(user *User, hub *Hub) error {
	return db.Model(user).Association("Hubs").Delete(hub)
}

func FollowUser(user *User, friend *User) error {
	return db.Model(user).Association("Friends").Append(friend)
}

func GetFollowingUsers(user *User) ([]*User, error) {
	var friends []*User
	err := db.Model(&user).Association("Friends").Find(&friends)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func UnfollowUser(user *User, friend *User) error {
	return db.Model(user).Association("Friends").Delete(friend)
}
