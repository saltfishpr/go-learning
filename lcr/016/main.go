package main

func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	l, r := 0, 0
	var res int
	for r < len(s) {
		rc := s[r]
		r++

		window[rc]++

		for window[rc] > 1 {
			lc := s[l]
			l++

			window[lc]--
		}

		res = max(res, r-l)
	}
	return res
}
