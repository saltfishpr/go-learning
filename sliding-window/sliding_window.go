// @file: sliding_window.go
// @date: 2021/1/16

// Package slidingwindow
package slidingwindow

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	// 目标状态
	var target = map[byte]int{}
	// 滑动窗口状态
	var window = map[byte]int{}
	// 这里不能用for range
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}
	// 滑动窗口 [left, right)
	left, right := 0, 0
	// 符合target的字母数量
	valid := 0
	// 最终结果字符串的起始及长度
	sLength := len(s)
	start, length := 0, sLength+1
	for right < sLength {
		// 取出字符
		var c = s[right]
		// 窗口扩容
		right++
		if target[c] != 0 {
			window[c]++
			if window[c] == target[c] {
				valid++
			}
		}
		// 缩小窗口的条件
		for valid == len(target) {
			// 更新最小子串
			if right-left < length {
				start = left
				length = right - left
			}
			// 取出窗口左侧的字符
			var d = s[left]
			// 缩小窗口
			left++
			if target[d] != 0 {
				if window[d] == target[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == sLength+1 {
		return ""
	}
	return s[start : start+length]
}

// 567. 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	var target = map[byte]int{}
	var window = map[byte]int{}
	for i := 0; i < len(s1); i++ {
		target[s1[i]]++
	}
	// 滑动窗口 [left, right)
	left, right := 0, 0
	// 符合target的字母数量
	valid := 0

	for right < len(s2) {
		c := s2[right]
		right++
		if target[c] != 0 {
			window[c]++
			if window[c] == target[c] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(target) {
				return true
			}
			d := s2[left]
			left++
			if target[d] != 0 {
				if window[d] == target[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

// 567. 字符串的排列 4ms
func checkInclusion2(s1 string, s2 string) bool {
	chars1 := []byte(s1)
	chars2 := []byte(s2)
	var c1 [26]int
	var c2 [26]int
	if len(s1) > len(s2) {
		return false
	}
	l := 0
	r := len(chars1)
	for i := range chars1 {
		c1[chars1[i]-'a']++
		c2[chars2[i]-'a']++
	}
	for r <= len(chars2) {
		if c1 == c2 {
			return true
		}
		if r < len(chars2) {
			c2[chars2[r]-'a']++
			c2[chars2[l]-'a']--
		}
		r++
		l++
	}
	return false
}

// 438. 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	res := []int{}
	if len(s) < len(p) {
		return res
	}
	chars := []byte(s)
	charp := []byte(p)
	var c1 [26]int
	var c2 [26]int
	// 记录字符个数
	for i := range p {
		c1[chars[i]-'a']++
		c2[charp[i]-'a']++
	}
	l, r := 0, len(p)
	for r <= len(s) {
		if c1 == c2 {
			res = append(res, l)
		}
		if r < len(s) {
			c1[chars[r]-'a']++
			c1[chars[l]-'a']--
		}
		l++
		r++
	}
	return res
}

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	var m = map[byte]int{}
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right++
		m[c]++
		for m[c] > 1 {
			d := s[left]
			left++
			m[d]--
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

// 30. 串联所有单词的子串
func findSubstring(s string, words []string) []int {
	n := len(s)
	wordNum := len(words)
	if n == 0 || wordNum == 0 {
		return []int{}
	}
	oneWord := len(words[0])
	if n < oneWord {
		return []int{}
	}

	targetCounter := make(map[string]int, 0)
	for _, w := range words {
		targetCounter[w]++
	}

	equal := func(a, b map[string]int) bool {
		if len(a) != len(b) {
			return false
		}
		for key := range a {
			if a[key] != b[key] {
				return false
			}
		}
		return true
	}

	res := make([]int, 0)
	// 对 0 ~ oneWord 循环，因为[0:]是[oneWord:]的子问题
	for i := 0; i < oneWord; i++ {
		curCnt := 0
		left, right := i, i
		curCounter := make(map[string]int, 0)
		for right+oneWord <= n {
			rightWord := s[right : right+oneWord]
			right += oneWord
			if _, ok := targetCounter[rightWord]; !ok {
				left = right
				curCnt = 0
				curCounter = make(map[string]int, 0)
			} else {
				curCounter[rightWord]++
				curCnt++
				for curCounter[rightWord] > targetCounter[rightWord] {
					leftWord := s[left : left+oneWord]
					left += oneWord
					curCounter[leftWord]--
					curCnt--
				}
				if equal(targetCounter, curCounter) {
					res = append(res, left)
				}
			}
		}
	}
	return res
}
