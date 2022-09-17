// @file: randomized_set.go
// @date: 2021/2/8

// Package randomizedset
package randomizedset

import "math/rand"

// 380. 常数时间插入、删除和获取随机元素
type RandomizedSet struct {
	valToIndex map[int]int

	nums []int
	size int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		valToIndex: make(map[int]int, 0),
		nums:       make([]int, 0),
		size:       0,
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.valToIndex[val]; ok {
		return false
	}
	s.nums = append(s.nums, val)
	s.valToIndex[val] = s.size
	s.size++
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.valToIndex[val]; !ok {
		return false
	}
	index := s.valToIndex[val]
	// 交换两个数
	s.nums[index], s.nums[s.size-1] = s.nums[s.size-1], s.nums[index]
	// 更新index
	s.valToIndex[s.nums[index]] = index
	// 删除val
	delete(s.valToIndex, val)
	s.nums = s.nums[:s.size-1]
	s.size--
	return true
}

/** Get a random element from the set. */
func (s *RandomizedSet) GetRandom() int {
	return s.nums[rand.Int()%s.size]
}
