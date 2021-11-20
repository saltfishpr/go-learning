// @description: 聊天室表字段
// @file: hub.go
// @date: 2021/11/19

package data

import (
	"gorm.io/plugin/soft_delete"
)

func init() {
	db := NewPostgres()
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
	db := NewPostgres()
	res := db.Create(hub)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func ReadAllHubs() ([]*Hub, error) {
	var hubs []*Hub

	db := NewPostgres()
	res := db.Find(&hubs)
	if res.Error != nil {
		return nil, res.Error
	}
	return hubs, nil
}

func UpdateHub(hub *Hub) error {
	db := NewPostgres()
	res := db.Model(hub).Where("h_id = ?", hub.HID).Updates(hub)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteHubByHID(hid string) error {
	db := NewPostgres()
	res := db.Where("h_id = ?", hid).Delete(&Hub{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func ReadHubByHID(hid string) (*Hub, error) {
	hub := new(Hub)

	db := NewPostgres()
	res := db.Where("h_id = ?", hid).First(hub)
	if res.Error != nil {
		return nil, res.Error
	}
	return hub, nil
}

func JoinHub(user *User, hub *Hub) error {
	db := NewPostgres()
	return db.Model(user).Association("Hubs").Append(hub)
}

func GetJoinedHubs(user *User) ([]*Hub, error) {
	var hubs []*Hub
	db := NewPostgres()
	err := db.Model(&user).Association("Hubs").Find(&hubs)
	if err != nil {
		return nil, err
	}
	return hubs, nil
}

func LeaveHub(user *User, hub *Hub) error {
	db := NewPostgres()
	return db.Model(user).Association("Hubs").Delete(hub)
}
