// @file: chatroom.go
// @date: 2021/11/10

package main

import (
	"log"
)

type ChatRoom struct{}

func ShowMessage(user User, message string) {
	log.Printf("%s: %s", user.Name, message)
}
