// @description: 提供数据库连接
// @file: connection.go
// @date: 2021/11/18

// Package data 提供数据库连接与数据存取操作.
package data

import (
	"fmt"

	"learning/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
	IsConnected() (bool, error)

	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUserByUsername(username string) error
	GetAllUsers() ([]*User, error)

	UserJoinHub(user *User, hub *Hub) error
	UserLeaveHub(user *User, hub *Hub) error
	GetUserJoinedHubs(user *User) ([]*Hub, error)
	UserFollowUser(user *User, followUser *User) error
	UserUnfollowUser(user *User, unfollowUser *User) error
	GetUserFollowingUsers(user *User) ([]*User, error)

	CreateHub(hub *Hub) error
	GetHubByHID(hid string) (*Hub, error)
	UpdateHub(hub *Hub) error
	DeleteHubByHID(hid string) error
	GetAllHubs() ([]*Hub, error)

	GetUsersInHub(hub *Hub) ([]*User, error)
	IsUserInHub(user *User, hub *Hub) bool
}

type Postgres struct {
	db *gorm.DB
}

var _ Connection = (*Postgres)(nil)

func NewPostgres() (Connection, error) {
	c := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		c.Postgres.Host,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Database,
		c.Postgres.Port,
	)
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: dsn,
			},
		),
	)
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func (p *Postgres) IsConnected() (bool, error) {
	db, err := p.db.DB()
	if err != nil {
		return false, err
	}
	return db.Ping() == nil, nil
}

func (p *Postgres) CreateUser(user *User) error {
	return p.db.Create(user).Error
}

func (p *Postgres) GetUserByUsername(username string) (*User, error) {
	var user User
	err := p.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *Postgres) UpdateUser(user *User) error {
	return p.db.Model(user).Where("username = ?", user.Username).Updates(user).Error
}

func (p *Postgres) DeleteUserByUsername(username string) error {
	return p.db.Where("username = ?", username).Delete(&User{}).Error
}

func (p *Postgres) GetAllUsers() ([]*User, error) {
	var users []*User
	err := p.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (p *Postgres) UserJoinHub(user *User, hub *Hub) error {
	return p.db.Model(user).Association("Hubs").Append(hub)
}

func (p *Postgres) UserLeaveHub(user *User, hub *Hub) error {
	return p.db.Model(user).Association("Hubs").Delete(hub)
}

func (p *Postgres) GetUserJoinedHubs(user *User) ([]*Hub, error) {
	var hubs []*Hub
	err := p.db.Model(user).Association("Hubs").Find(&hubs)
	if err != nil {
		return nil, err
	}
	return hubs, nil
}

func (p *Postgres) UserFollowUser(user *User, followUser *User) error {
	return p.db.Model(user).Association("Friends").Append(followUser)
}

func (p *Postgres) UserUnfollowUser(user *User, unfollowUser *User) error {
	return p.db.Model(user).Association("Friends").Delete(unfollowUser)
}

func (p *Postgres) GetUserFollowingUsers(user *User) ([]*User, error) {
	var users []*User
	err := p.db.Model(user).Association("Friends").Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (p *Postgres) CreateHub(hub *Hub) error {
	return p.db.Create(hub).Error
}

func (p *Postgres) GetHubByHID(hid string) (*Hub, error) {
	var hub Hub
	err := p.db.Where("h_id = ?", hid).First(&hub).Error
	if err != nil {
		return nil, err
	}
	return &hub, nil
}

func (p *Postgres) UpdateHub(hub *Hub) error {
	return p.db.Model(hub).Where("h_id = ?", hub.HID).Updates(hub).Error
}

func (p *Postgres) DeleteHubByHID(hid string) error {
	return p.db.Where("h_id = ?", hid).Delete(&Hub{}).Error
}

func (p *Postgres) GetAllHubs() ([]*Hub, error) {
	var hubs []*Hub
	err := p.db.Find(&hubs).Error
	if err != nil {
		return nil, err
	}
	return hubs, nil
}

func (p *Postgres) GetUsersInHub(hub *Hub) ([]*User, error) {
	var users []*User
	err := p.db.Model(hub).Association("Users").Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (p *Postgres) IsUserInHub(user *User, hub *Hub) bool {
	return p.db.Model(user).Association("Hubs").Find(hub) == nil
}
