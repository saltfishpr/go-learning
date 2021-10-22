// @file: socket.go
// @description: 使用 net 包从 socket 中打开，写入，读取数据
// @author: SaltFish
// @date: 2020/09/08

// Package ch15 is chapter 15
package ch15

import (
	"fmt"
	"io"
	"net"
)

func MySocket() {
	var (
		host   = "www.apache.org"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 4096)
		read   = true
		count  = 0
	)
	// 创建一个socket
	con, err := net.Dial("tcp", remote)
	// 发送我们的消息，一个http GET请求
	io.WriteString(con, msg)
	// 读取服务器的响应
	for read {
		count, err = con.Read(data)
		read = err == nil
		fmt.Printf(string(data[0:count]))
	}
	con.Close()
}
