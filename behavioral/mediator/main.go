// @file: main.go
// @date: 2021/11/10

package main

func main() {
	robert := &User{Name: "Robert"}
	john := &User{Name: "John"}

	robert.SendMessage("Hi! John!")
	john.SendMessage("Hello! Robert!")
}
