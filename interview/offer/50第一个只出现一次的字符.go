// @file: 50第一个只出现一次的字符.go
// @date: 2021/2/21

// Package offer
package offer

/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
*/

func firstUniqCharX50(s string) byte {
	counter := make([]int, 256)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if counter[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
