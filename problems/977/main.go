package main

func sortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	for k := len(nums) - 1; k >= 0; k-- {
		a, b := nums[i]*nums[i], nums[j]*nums[j]
		if a > b {
			res[k] = a
			i++
		} else {
			res[k] = b
			j--
		}
	}
	return res
}
