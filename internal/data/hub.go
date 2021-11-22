// @description: 聊天室表字段
// @file: hub.go
// @date: 2021/11/19

package data

import (
	"gorm.io/plugin/soft_delete"
)

func init() {
	_ = NewPostgres()
	db.AutoMigrate(&Hub{})
}

type Hub struct {
	Base
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:hub_udx_delete"`

	HID  *string `gorm:"size:8;not null;uniqueIndex:hub_udx_delete"`
	Name *string `gorm:"not null"`
	Size *int    `gorm:"not null"`

	Users []*User `gorm:"many2many:user_hub"`
}

func CreateHub(hub *Hub) error {
	res := db.Create(hub)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetAllHubs() ([]*Hub, error) {
	var hubs []*Hub

	res := db.Find(&hubs)
	if res.Error != nil {
		return nil, res.Error
	}
	return hubs, nil
}

func UpdateHub(hub *Hub) error {
	res := db.Model(hub).Where("h_id = ?", hub.HID).Updates(hub)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteHubByHID(hid string) error {
	res := db.Where("h_id = ?", hid).Delete(&Hub{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetHubByHID(hid string) (*Hub, error) {
	hub := new(Hub)

	res := db.Where("h_id = ?", hid).First(hub)
	if res.Error != nil {
		return nil, res.Error
	}
	return hub, nil
}

func GetUsersInHub(hub *Hub) ([]*User, error) {
	var users []*User
	err := db.Model(&hub).Association("Users").Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func IsInHub(hub *Hub, user *User) (bool, error) {
	var res = new(User)
	err := db.Model(&hub).Where("user_id = ?", user.ID).Association("Users").Find(&res)
	if err != nil {
		return false, err
	}
	return res != nil, nil
}
