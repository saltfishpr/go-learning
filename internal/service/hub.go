// @description:
// @file: hub.go
// @date: 2021/11/22

package service

import (
	"github.com/jinzhu/copier"

	"learning/internal/data"
	"learning/internal/model"
)

func CreateHub(hub *model.Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)
	return data.CreateHub(hubEntity)
}

func GetAllHubs() ([]*model.Hub, error) {
	hubEntities, err := data.GetAllHubs()
	if err != nil {
		return nil, err
	}
	hubs := make([]*model.Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)
	return hubs, nil
}

func UpdateHub(hub *model.Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)
	return data.UpdateHub(hubEntity)
}

func DeleteHubByHID(hid string) error {
	return data.DeleteHubByHID(hid)
}

func GetHubByHID(hid string) (*model.Hub, error) {
	hubEntity, err := data.GetHubByHID(hid)
	if err != nil {
		return nil, err
	}
	hub := new(model.Hub)
	copier.Copy(hub, hubEntity)

	return hub, nil
}
