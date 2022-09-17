// @file: blacklist_random.go
// @date: 2021/2/8

// Package blacklistrandom
package blacklistrandom

import "math/rand"

// 710. 黑名单中的随机数
type Solution struct {
	nums []int
}

func Constructor(N int, blacklist []int) Solution {
	blackMap := make(map[int]struct{})
	for _, val := range blacklist {
		blackMap[val] = struct{}{}
	}
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i
	}
	last := N - 1
	tail := last - len(blacklist)
	for _, val := range blacklist {
		if val > tail {
			continue
		}
		for _, ok := blackMap[last]; ok; _, ok = blackMap[last] {
			last--
		}
		nums[val] = nums[last]
		last--
	}
	return Solution{nums: nums[:tail+1]}
}

func (s *Solution) Pick() int {
	return s.nums[rand.Int()%len(s.nums)]
}
