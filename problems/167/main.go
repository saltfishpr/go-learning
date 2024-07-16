package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		current := numbers[l] + numbers[r]
		if current == target {
			return []int{l + 1, r + 1}
		} else if current < target {
			l++
		} else {
			r--
		}
	}
	return []int{-1, -1}
}
