// @description: 用户表字段
// @file: user.go
// @date: 2021/11/19

package data

import (
	"gorm.io/plugin/soft_delete"
)

func init() {
	db := NewPostgres()
	db.AutoMigrate(&User{})
}

type User struct {
	Base
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:user_udx_delete"`

	Account  *string `gorm:"size:32;not null;uniqueIndex:user_udx_delete"`
	Password *string `gorm:"size:32;not null"`
	Nickname *string `gorm:"size:64;not null"`
	Address  *string `gorm:"size:128"`

	Hubs    []*Hub  `gorm:"many2many:user_hub"`
	Friends []*User `gorm:"many2many:user_friends"`
}

func CreateUser(user *User) error {
	db := NewPostgres()
	res := db.Create(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func ReadAllUsers() ([]*User, error) {
	var users []*User

	db := NewPostgres()
	res := db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func UpdateUser(user *User) error {
	db := NewPostgres()
	res := db.Model(user).Where("account = ?", user.Account).Updates(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteUserByAccount(account string) error {
	db := NewPostgres()
	res := db.Where("account = ?", account).Delete(&User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func ReadUserByAccount(account string) (*User, error) {
	user := new(User)

	db := NewPostgres()
	res := db.Where("account = ?", account).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
