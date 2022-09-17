// @file: user.go
// @date: 2021/11/10

package main

type User struct {
	Name string
}

func (u User) SendMessage(message string) {
	ShowMessage(u, message)
}
