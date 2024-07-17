package main

func findAnagrams(s string, p string) []int {
	need := map[byte]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	window := map[byte]int{}
	var valid int
	var res []int
	l, r := 0, 0
	for r < len(s) {
		rc := s[r]
		r++

		if _, ok := need[rc]; ok {
			window[rc]++
			if window[rc] == need[rc] {
				valid++
			}
		}

		for r-l >= len(p) {
			if valid == len(need) {
				res = append(res, l)
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

	return res
}
