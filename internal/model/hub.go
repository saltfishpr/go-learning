// @description: 聊天室数据结构
// @file: model.go
// @date: 2021/11/1

package model

import (
	"github.com/jinzhu/copier"

	"learning/internal/data"
)

type Hub struct {
	HID  *string `json:"hid" validate:"required"`
	Name *string `json:"name" validate:"required"`
	Size *int    `json:"size" validate:"required"`

	Users []*User `json:"users,omitempty"`
}

func CreateHub(hub *Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)
	return data.CreateHub(hubEntity)
}

func ReadAllHubs() ([]*Hub, error) {
	hubEntities, err := data.ReadAllHubs()
	if err != nil {
		return nil, err
	}
	hubs := make([]*Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)
	return hubs, nil
}

func UpdateHub(hub *Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)
	return data.UpdateHub(hubEntity)
}

func DeleteHubByHID(hid string) error {
	return data.DeleteHubByHID(hid)
}

func JoinHub(account string, hid string) error {
	userEntity, err := data.ReadUserByAccount(account)
	if err != nil {
		return err
	}

	hubEntity, err := data.ReadHubByHID(hid)
	if err != nil {
		return err
	}

	return data.JoinHub(userEntity, hubEntity)
}

func ReadHubByHID(hid string) (*Hub, error) {
	hubEntity, err := data.ReadHubByHID(hid)
	if err != nil {
		return nil, err
	}
	hub := new(Hub)
	copier.Copy(hub, hubEntity)
	return hub, nil
}

func LeaveHub(account string, hid string) error {
	userEntity, err := data.ReadUserByAccount(account)
	if err != nil {
		return err
	}

	hubEntity, err := data.ReadHubByHID(hid)
	if err != nil {
		return err
	}

	return data.LeaveHub(userEntity, hubEntity)
}
