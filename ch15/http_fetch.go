// @file: http_fetch.go
// @description: 使用 http.Get() 获取并显示网页内容
// @author: SaltFish
// @date: 2020/09/08

// Package ch15 is chapter 15
package ch15

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func MyFetch() {
	res, err := http.Get("http://www.google.com")
	checkError2(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError2(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError2(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
