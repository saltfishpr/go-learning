package main

func carPooling(trips [][]int, capacity int) bool {
	n := 1001 // 最多有 1000 个站
	diff := make([]int, n)
	for _, trip := range trips {
		num, from, to := trip[0], trip[1], trip[2]
		diff[from] += num
		diff[to] -= num
	}
	if diff[0] > capacity {
		return false
	}
	for i := 1; i < n; i++ {
		diff[i] = diff[i] + diff[i-1]
		if diff[i] > capacity {
			return false
		}
	}
	return true
}
