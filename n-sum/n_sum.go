// @file: n_sum.go
// @date: 2021/1/19

// Package nsum
package nsum

import "sort"

// n 递归数目，start 左边界
func nSum(nums []int, n, start, target int) [][]int {
	length := len(nums)
	var res [][]int
	if length < n {
		return res
	}

	// two sum 双指针技巧
	if n == 2 {
		l, r := start, length-1
		for l < r {
			numL, numR := nums[l], nums[r]
			sum := numL + numR
			if sum < target {
				// 优化重复元素的计算
				for l < r && nums[l] == numL {
					l++
				}
			} else if sum > target {
				for l < r && nums[r] == numR {
					r--
				}
			} else {
				res = append(res, []int{numL, numR})
				for l < r && nums[l] == numL {
					l++
				}
				for l < r && nums[r] == numR {
					r--
				}
			}
		}
		return res
	}

	for i := start; i < length; i++ {
		// 递归计算 target-nums[i] 的所有可能组合
		tuples := nSum(nums, n-1, i+1, target-nums[i])
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}
		// 优化重复元素的计算
		for i < length-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]] = i
	}
	for i := range nums {
		need := target - nums[i]
		if v, ok := m[need]; ok && v != i {
			return []int{i, v}
		}
	}
	return []int{-1, -1}
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n <= 2 {
		return res
	}
	sort.Ints(nums)

	twoS := func(nums []int, start int, target int) [][]int {
		l, r := start, len(nums)-1
		res := [][]int{}

		for l < r {
			sum := nums[l] + nums[r]
			num1, num2 := nums[l], nums[r]
			if sum < target {
				for l < r && nums[l] == num1 {
					l++
				}
			} else if sum > target {
				for l < r && nums[r] == num2 {
					r--
				}
			} else {
				res = append(res, []int{num1, num2})
				for l < r && nums[l] == num1 {
					l++
				}
				for l < r && nums[r] == num2 {
					r--
				}
			}
		}
		return res
	}

	for i := 0; i < n; i++ {
		tmp := twoS(nums, i+1, -nums[i])
		for _, v := range tmp {
			res = append(res, append(v, nums[i]))
		}
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// 18. 四数之和
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 4, 0, target)
}
