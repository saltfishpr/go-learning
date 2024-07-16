package main

import "math/rand"

type RandomizedSet struct {
	data []int
	m    map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		data: []int{},
		m:    map[int]int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.data)
	this.data = append(this.data, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.m[val]
	if !ok {
		return false
	}
	lastVal := this.data[len(this.data)-1]
	this.data[idx] = lastVal
	this.m[lastVal] = idx

	delete(this.m, val)
	this.data = this.data[:len(this.data)-1]

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.data))
	return this.data[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
