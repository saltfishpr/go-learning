package main

import (
	"crypto/rand"
)

func main() {
	s, err := generate(7)
	if err != nil {
		panic(err)
	}
	println(s)
}

// generate 生成随机的 base62 字符串.
func generate(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		bytes[i] = chars[bytes[i]%byte(len(chars))]
	}
	return string(bytes), nil
}
