// @file: 38字符串的排列.go
// @date: 2021/2/18

// Package offer
package offer

/*
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
*/

func permutationX38(s string) []string {
	set := make(map[string]struct{}, 0)
	res := make([]string, 0)
	n := len(s)
	counter := make(map[byte]int, 0)
	for i := 0; i < n; i++ {
		counter[s[i]]++
	}

	var helper func([]byte)
	helper = func(bytes []byte) {
		if len(bytes) == n {
			ss := string(bytes)
			if _, ok := set[ss]; !ok {
				res = append(res, ss)
				set[ss] = struct{}{}
			}
			return
		}

		for i := 0; i < n; i++ {
			c := s[i]
			if counter[c] != 0 {
				bytes = append(bytes, c)
				counter[c]--
				helper(bytes)
				bytes = bytes[:len(bytes)-1]
				counter[c]++
			}
		}

	}

	helper([]byte{})
	return res
}

func permutationX38V2(s string) []string {
	bytes := []byte(s)
	res := make([]string, 0)
	n := len(s)

	var helper func(int)
	helper = func(k int) {
		if k == n-1 {
			res = append(res, string(bytes))
		}
		// 第 k 位使用过的字符
		set := make([]bool, 256)
		for i := k; i < n; i++ {
			// 剪枝
			if set[bytes[i]] {
				continue
			}
			set[bytes[i]] = true
			// 将第 i 位放到第 k 位上
			bytes[i], bytes[k] = bytes[k], bytes[i]
			// 递归第 k+1 位字符
			helper(k + 1)
			// 回溯
			bytes[i], bytes[k] = bytes[k], bytes[i]
		}
	}

	helper(0)
	return res
}
