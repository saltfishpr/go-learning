// @description: 保存websocket连接
// @file: storage.go
// @date: 2021/11/21

package connstorage

import (
	"github.com/gofiber/websocket/v2"
	cmap "github.com/orcaman/concurrent-map"
)

var switchboard = cmap.New()

func Set(sid string, conn *websocket.Conn) {
	switchboard.Set(sid, conn)
}

func Get(sid string) (*websocket.Conn, bool) {
	if conn, ok := switchboard.Get(sid); ok {
		return conn.(*websocket.Conn), true
	}
	return nil, false
}

func Del(sid string) {
	switchboard.Remove(sid)
}
