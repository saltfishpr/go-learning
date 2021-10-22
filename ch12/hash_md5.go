// @file: hash_md5.go
// @description:
// @author: SaltFish
// @date: 2020/08/31

// Package ch12 is chapter 12
package ch12

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MyMd5() {
	hasher := md5.New()
	b := []byte{}
	io.WriteString(hasher, "test")
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
}
