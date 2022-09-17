// @file: search.go
// @date: 2021/1/16

// Package binarysearch
package binarysearch

import (
	"math"
)

// 找到target在数组中的位置
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

// 找到数组nums中target的左侧边界
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	// 判断是否越界
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 找到数组nums中target的右侧边界
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// 704. 二分查找
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

// 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return left
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix)-1, len(matrix[0])-1
	rowTop, rowBottom := 0, rows
	for rowTop <= rowBottom {
		rowMid := rowTop + (rowBottom-rowTop)/2
		if target >= matrix[rowMid][0] && target <= matrix[rowMid][cols] {
			rowTop = rowMid
			break // target在rowMid这一行
		} else if target > matrix[rowMid][cols] {
			rowTop = rowMid + 1
		} else if target < matrix[rowMid][0] {
			rowBottom = rowMid - 1
		}
	}
	if rowTop < 0 || rowTop > rows {
		return false
	}

	nums := matrix[rowTop]
	left, right := 0, cols
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return true
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return false
}

// 278. 第一个错误的版本
func isBadVersion(version int) bool {
	if version == 0 {
		return false
	}
	return true
}

func firstBadVersion(n int) int {
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// 154. 寻找旋转排序数组中的最小值 II
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

// 33. 搜索旋转排序数组
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 旋转点在[mid, right]中
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func search2(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		// 旋转点在[mid, right]中
		if nums[left] == nums[mid] {
			left++
		} else if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

// 875. 爱吃香蕉的珂珂
func minEatingSpeed(piles []int, H int) int {
	max := func(nums []int) int {
		maxVal := math.MinInt64
		for i := range nums {
			if nums[i] > maxVal {
				maxVal = nums[i]
			}
		}
		return maxVal
	}

	canFinish := func(nums []int, k int, H int) bool {
		times := 0
		for i := range nums {
			if nums[i]%k == 0 {
				times += nums[i] / k
			} else {
				times += nums[i]/k + 1
			}
		}
		return times <= H
	}
	n := max(piles)
	// 速度最小值为1， 最大值为max(piles)
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if !canFinish(piles, mid, H) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 1011. 在 D 天内送达包裹的能力
func shipWithinDays(weights []int, D int) int {
	maxAndSum := func(nums []int) (int, int) {
		maxVal := math.MinInt64
		sum := 0
		for i := range nums {
			if nums[i] > maxVal {
				maxVal = nums[i]
			}
			sum += nums[i]
		}
		return maxVal, sum
	}

	canFinish := func(nums []int, k int, D int) bool {
		days := 1
		weight := 0
		for i := range nums {
			if weight+nums[i] <= k {
				weight += nums[i]
			} else {
				days++
				weight = nums[i]
			}
		}
		return days <= D
	}

	n, sum := maxAndSum(weights)
	// 速度最小值为1， 最大值为max(piles)
	left, right := n, sum
	for left <= right {
		mid := left + (right-left)/2
		if !canFinish(weights, mid, D) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 410. 分割数组的最大值
func splitArray(nums []int, m int) int {
	maxAndSum := func(nums []int) (int, int) {
		max, sum := math.MinInt64, 0
		for i := range nums {
			if nums[i] > max {
				max = nums[i]
			}
			sum += nums[i]
		}
		return max, sum
	}

	split := func(nums []int, max int) int {
		count := 1
		weight := 0
		for i := range nums {
			if weight+nums[i] <= max {
				weight += nums[i]
			} else {
				count++
				weight = nums[i]
			}
		}
		return count
	}

	max, sum := maxAndSum(nums)
	left, right := max, sum
	for left < right {
		mid := left + (right-left)/2
		if split(nums, mid) <= m {
			// 需要的分割次数少，说明mid太高
			right = mid
		} else if split(nums, mid) > m {
			// 需要的分割次数多，说明mid太低
			left = mid + 1
		}
	}
	return left
}
