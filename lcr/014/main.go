package main

func checkInclusion(s1 string, s2 string) bool {
	need := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	window := map[byte]int{}
	l, r := 0, 0
	valid := 0
	for r < len(s2) {
		rc := s2[r]
		r++

		if _, ok := need[rc]; ok {
			window[rc]++
			if window[rc] == need[rc] {
				valid++
			}
		}

		for r-l >= len(s1) {
			if valid == len(need) {
				return true
			}

			lc := s2[l]
			l++

			if _, ok := need[lc]; ok {
				if window[lc] == need[lc] {
					valid--
				}
				window[lc]--
			}
		}
	}

	return false
}
