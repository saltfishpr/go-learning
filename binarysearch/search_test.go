// @file: search_test.go
// @date: 2021/1/26

// Package binarysearch
package binarysearch

import "testing"

func TestSearchMatrix(t *testing.T) {
	data := [][]int{{1}}
	result := searchMatrix(data, 2)
	if result != false {
		t.Fatalf("get:%v, want false", result)
	}
}
