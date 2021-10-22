// @file: 58-1翻转单词顺序.go
// @date: 2021/2/24

// Package offer
package offer

import "strings"

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
*/

func reverseWordsX58(s string) string {
	reverse := func(strs []string) {
		l, r := 0, len(strs)-1
		for l < r {
			strs[l], strs[r] = strs[r], strs[l]
			l++
			r--
		}
	}

	strs := strings.Split(s, " ")
	reverse(strs)
	var builder strings.Builder
	for i := 0; i < len(strs); i++ {
		if strs[i] == "" {
			continue
		}
		builder.WriteString(" ")
		builder.WriteString(strs[i])
	}
	res := builder.String()
	if len(res) == 0 {
		return ""
	}
	return res[1:]
}
