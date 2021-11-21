// @description: 保存websocket连接
// @file: conn_storage.go
// @date: 2021/11/21

package utils

import (
	"github.com/gofiber/websocket/v2"
	cmap "github.com/orcaman/concurrent-map"
)

var switchboard = cmap.New()

func Connect(account string, conn *websocket.Conn) {
	switchboard.SetIfAbsent(account, conn)
}

func GetConnection(account string) *websocket.Conn {
	if conn, ok := switchboard.Get(account); ok {
		return conn.(*websocket.Conn)
	}
	return nil
}

func Disconnect(account string) {
	switchboard.Remove(account)
}
