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
}

func CreateHub(hub *Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)
	return data.CreateHub(hubEntity)
}

func GetAllHubs() ([]*Hub, error) {
	hubEntities, err := data.GetAllHubs()
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

func GetHubByHID(hid string) (*Hub, error) {
	hubEntity, err := data.GetHubByHID(hid)
	if err != nil {
		return nil, err
	}
	hub := new(Hub)
	copier.Copy(hub, hubEntity)

	return hub, nil
}
