// @file: 58-2左旋转字符串.go
// @date: 2021/2/24

// Package offer
package offer

import "strings"

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
*/

func reverseLeftWordsX58(s string, n int) string {
	left, right := s[:n], s[n:]
	var builder strings.Builder
	builder.WriteString(right)
	builder.WriteString(left)
	return builder.String()
}
