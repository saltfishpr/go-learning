package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var root1 = &TreeNode{
	Val: 10,
	Left: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: -2},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 1},
		},
	},
	Right: &TreeNode{
		Val:   -3,
		Right: &TreeNode{Val: 11},
	},
}

var root2 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   11,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 2},
		},
	},
	Right: &TreeNode{
		Val:  8,
		Left: &TreeNode{Val: 13},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1},
		},
	},
}

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				root:      root1,
				targetSum: 8,
			},
			want: 3,
		},
		{
			name: "示例 2",
			args: args{
				root:      root2,
				targetSum: 22,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pathSum(tt.args.root, tt.args.targetSum))
		})
	}
}

var result int

func Benchmark_pathSum(b *testing.B) {
	var res int
	for n := 0; n < b.N; n++ {
		res = pathSum(root2, 22)
	}
	result = res
}
