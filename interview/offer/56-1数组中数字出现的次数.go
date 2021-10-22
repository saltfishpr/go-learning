// @file: 56-1数组中数字出现的次数.go
// @date: 2021/2/23

// Package offer
package offer

/*
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
*/

func singleNumbersX56(nums []int) []int {
	// mask 为a, b两个数字异或的结果
	mask := 0
	for _, num := range nums {
		mask ^= num
	}
	// 两个数字不同的一个二进制位
	diff := (-mask) & mask
	x := 0
	for _, num := range nums {
		if diff&num != 0 {
			x ^= num
		}
	}
	// mask = a ^ b, 所以 mask ^ a = b
	return []int{x, x ^ mask}
}
