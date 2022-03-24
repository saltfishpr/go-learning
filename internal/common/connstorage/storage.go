// @file: storage.go
// @date: 2021/11/21

// Package connstorage 保存websocket连接.
package connstorage

import (
	"github.com/gofiber/websocket/v2"
	cmap "github.com/orcaman/concurrent-map"
)

var switchboard = cmap.New()

func Set(username string, conn *websocket.Conn) {
	switchboard.Set(username, conn)
}

func Get(username string) (*websocket.Conn, bool) {
	if conn, ok := switchboard.Get(username); ok {
		return conn.(*websocket.Conn), true
	}
	return nil, false
}

func Del(username string) {
	switchboard.Remove(username)
}
