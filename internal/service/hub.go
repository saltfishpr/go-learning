// @description: 聊天室服务
// @file: hub.go
// @date: 2021/11/22

package service

import (
	"learning/internal/data"
	"learning/internal/model"

	"github.com/jinzhu/copier"
)

type IHub interface {
	CreateHub(hub *model.Hub) error
	GetHubByHID(hid string) (*model.Hub, error)
	UpdateHub(hub *model.Hub) error
	DeleteHubByHID(hid string) error
	GetAllHubs() ([]*model.Hub, error)
}

type Hub struct {
	connection data.Connection
}

var _ IHub = (*Hub)(nil)

func NewHub(connection data.Connection) *Hub {
	return &Hub{
		connection: connection,
	}
}

func (h *Hub) CreateHub(hub *model.Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)

	return h.connection.CreateHub(hubEntity)
}

func (h *Hub) GetHubByHID(hid string) (*model.Hub, error) {
	hubEntity, err := h.connection.GetHubByHID(hid)
	if err != nil {
		return nil, err
	}
	hub := new(model.Hub)
	copier.Copy(hub, hubEntity)

	return hub, nil
}

func (h *Hub) UpdateHub(hub *model.Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)

	return h.connection.UpdateHub(hubEntity)
}

func (h *Hub) DeleteHubByHID(hid string) error {
	return h.connection.DeleteHubByHID(hid)
}

func (h *Hub) GetAllHubs() ([]*model.Hub, error) {
	hubEntities, err := h.connection.GetAllHubs()
	if err != nil {
		return nil, err
	}
	hubs := make([]*model.Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)

	return hubs, nil
}
