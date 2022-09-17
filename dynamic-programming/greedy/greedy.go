// @file: greedy.go
// @date: 2021/3/8

// Package greedy
package greedy

import "sort"

// 435. 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(
		intervals, func(i, j int) bool {
			a, b := intervals[i], intervals[j]
			return a[1] < b[1]
		},
	)

	res := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			res++
			continue
		}
		end = intervals[i][1]
	}
	return res
}

// 452. 用最少数量的箭引爆气球
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(
		points, func(i, j int) bool {
			a, b := points[i], points[j]
			return a[1] < b[1]
		},
	)

	res := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if end > points[i][0] {
			continue
		}
		res++
		end = points[i][1]
	}
	return res
}

// 55. 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	farthest := 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest < i+1 {
			return false
		}
	}
	return farthest >= n-1
}

// 45. 跳跃游戏 II
func jump(nums []int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := 0
	end := 0
	farthest := 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, nums[i]+i)
		if end == i {
			res++
			end = farthest
		}
	}
	return res
}
