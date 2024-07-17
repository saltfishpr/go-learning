package main

func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	var res int
	l, r := 0, 0
	for r < len(s) {
		c := s[r]
		window[c]++
		r++

		for window[c] > 1 {
			window[s[l]]--
			l++
		}

		res = max(res, r-l)
	}
	return res
}
