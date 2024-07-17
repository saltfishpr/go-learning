package main

func minWindow(s string, t string) string {
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := map[byte]int{}
	valid := 0
	l, r := 0, 0
	start, minLength := 0, len(s)+1
	for r < len(s) {
		rc := s[r]
		r++

		if _, ok := need[rc]; ok {
			window[rc]++
			if window[rc] == need[rc] {
				valid++
			}
		}

		for valid == len(need) {
			if r-l < minLength {
				start = l
				minLength = r - l
			}

			lc := s[l]
			l++

			if _, ok := need[lc]; ok {
				if window[lc] == need[lc] {
					valid--
				}
				window[lc]--
			}
		}
	}

	if minLength == len(s)+1 {
		return ""
	}
	return s[start : start+minLength]
}
