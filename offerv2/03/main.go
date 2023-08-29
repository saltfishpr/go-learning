package main

func findRepeatNumber(nums []int) int {
	storage := make([]int, len(nums))
	for _, num := range nums {
		storage[num]++
		if storage[num] > 1 {
			return num
		}
	}
	return 0
}
