// @file: server.go
// @description: 网络，模板和网页应用
// @author: SaltFish
// @date: 2020/09/08

// Package ch15 is chapter 15
package ch15

import (
	"fmt"
	"net"
)

func MyServer() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // 终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
