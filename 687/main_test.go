package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestUnivaluePath(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 1},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			want: 2,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestUnivaluePath(tt.args.root))
		})
	}
}

var result int

func Benchmark_longestUnivaluePath(b *testing.B) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 5},
		},
	}

	var res int
	for i := 0; i < b.N; i++ {
		res = longestUnivaluePath(root)
	}
	result = res
}
