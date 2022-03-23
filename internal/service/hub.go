// @description: 聊天室服务
// @file: hub.go
// @date: 2021/11/22

package service

import (
	"learning/internal/data"
	"learning/internal/model"

	"github.com/jinzhu/copier"
)

type Hub struct {
	conn data.Connection
}

func NewHub(conn data.Connection) *Hub {
	return &Hub{
		conn: conn,
	}
}

func (h *Hub) CreateHub(hub *model.Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)
	return h.conn.CreateHub(hubEntity)
}

func (h *Hub) GetAllHubs() ([]*model.Hub, error) {
	hubEntities, err := h.conn.GetAllHubs()
	if err != nil {
		return nil, err
	}
	hubs := make([]*model.Hub, len(hubEntities))
	copier.Copy(&hubs, &hubEntities)
	return hubs, nil
}

func (h *Hub) UpdateHub(hub *model.Hub) error {
	hubEntity := new(data.Hub)
	copier.Copy(hubEntity, hub)
	return h.conn.UpdateHub(hubEntity)
}

func (h *Hub) DeleteHubByHID(hid string) error {
	return h.conn.DeleteHubByHID(hid)
}

func (h *Hub) GetHubByHID(hid string) (*model.Hub, error) {
	hubEntity, err := h.conn.GetHubByHID(hid)
	if err != nil {
		return nil, err
	}
	hub := new(model.Hub)
	copier.Copy(hub, hubEntity)

	return hub, nil
}
