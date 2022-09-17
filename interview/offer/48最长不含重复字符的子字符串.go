// @file: 48最长不含重复字符的子字符串.go
// @date: 2021/2/21

// Package offer
package offer

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
*/

func lengthOfLongestSubstringX48(s string) int {
	res := 0
	i := -1
	cToIdx := make(map[byte]int, 0)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for j := 0; j < len(s); j++ {
		c := s[j]
		if _, ok := cToIdx[c]; ok {
			i = max(cToIdx[c], i) // 更新左指针
		}
		cToIdx[c] = j
		res = max(res, j-i) // 更新结果
	}
	return res
}
