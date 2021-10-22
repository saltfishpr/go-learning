// @file: 56-2数组中数字出现的次数.go
// @date: 2021/2/23

// Package offer
package offer

/*
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
*/

func singleNumberX56(nums []int) int {
	binary := make([]int, 64)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := 0; j < 64; j++ {
			binary[j] += num & 1
			num >>= 1
		}
	}
	res := 0
	x := 1
	for i := 0; i < 64; i++ {
		res += x * (binary[i] % 3)
		x *= 2
	}
	return res
}

func singleNumberX56V2(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}
