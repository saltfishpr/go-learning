// @description: 用户数据结构与持久化
// @file: user.go
// @date: 2021/11/18

package model

import (
	"github.com/jinzhu/copier"

	"learning/internal/data"
)

func init() {
	db := data.NewPostgres()
	db.AutoMigrate(&UserEntity{})
}

type UserEntity struct {
	Base
	User
}

type User struct {
	Account  *string `json:"account" gorm:"size:32;not null;uniqueIndex:udx_delete" validate:"required"`
	Password *string `json:"password" gorm:"size:32;not null"`
	Nickname *string `json:"nickname" gorm:"size:64;not null"`
	Address  *string `json:"address,omitempty"`
}

func CreateUser(user *User) error {
	userEntity := new(UserEntity)
	copier.Copy(userEntity, user)

	db := data.NewPostgres()
	res := db.Create(userEntity)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func ReadUserByAccount(account string) (*User, error) {
	user := new(User)
	userEntity := new(UserEntity)

	db := data.NewPostgres()
	res := db.Where("account = ?", account).First(userEntity)
	if res.Error != nil {
		return nil, res.Error
	}
	copier.Copy(user, userEntity)
	return user, nil
}

func ReadAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	userEntities := make([]*UserEntity, 0)

	db := data.NewPostgres()
	res := db.Find(&userEntities)
	if res.Error != nil {
		return nil, res.Error
	}
	copier.Copy(&users, &userEntities)
	return users, nil
}

func UpdateUser(user *User) error {
	userEntity := new(UserEntity)
	copier.Copy(userEntity, user)

	db := data.NewPostgres()
	res := db.Model(userEntity).Where("account = ?", userEntity.Account).Updates(userEntity)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteUserByAccount(account string) error {
	db := data.NewPostgres()
	res := db.Where("account = ?", account).Delete(&UserEntity{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
