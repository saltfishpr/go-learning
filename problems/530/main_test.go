package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMinimumDifference(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 6},
				},
			},
			want: 1,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   48,
						Left:  &TreeNode{Val: 12},
						Right: &TreeNode{Val: 49},
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getMinimumDifference(tt.args.root))
		})
	}
}

var result int

func Benchmark_getMinimumDifference(b *testing.B) {
	var res int

	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}

	for n := 0; n < b.N; n++ {
		res = getMinimumDifference(root)
	}

	result = res
}
