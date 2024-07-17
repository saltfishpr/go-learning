package main

func dismantlingAction(arr string) int {
	window := map[byte]int{}
	l, r := 0, 0
	res := 0
	for r < len(arr) {
		rc := arr[r]
		r++

		window[rc]++

		for window[rc] > 1 {
			lc := arr[l]
			l++

			window[lc]--
		}

		res = max(res, r-l)
	}
	return res
}
