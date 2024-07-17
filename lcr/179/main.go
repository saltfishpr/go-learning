package main

func twoSum(price []int, target int) []int {
	l, r := 0, len(price)-1
	for l < r {
		sum := price[l] + price[r]
		if sum == target {
			return []int{price[l], price[r]}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
