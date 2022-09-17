// @file: 05替换空格.go
// @date: 2021/2/12

// Package offer
package offer

import "strings"

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/

func replaceSpaceX05(s string) string {
	bs := []byte(s)
	var builder strings.Builder
	for _, c := range bs {
		if c == ' ' {
			builder.Write([]byte{'%', '2', '0'})
		} else {
			builder.WriteByte(c)
		}
	}
	return builder.String()
}
