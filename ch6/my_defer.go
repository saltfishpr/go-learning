// @file: my_defer.go
// @description: defer语句在函数返回后执行语句
// @author: SaltFish
// @date: 2020/08/04

// Package ch6 is chapter 6
package ch6

import "fmt"

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() // function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return // terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}

// MyDefer is fun
func MyDefer() {
	doDBOperations()
}
