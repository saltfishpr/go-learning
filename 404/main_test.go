package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumOfLeftLeaves(t *testing.T) {
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
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: 24,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sumOfLeftLeaves(tt.args.root))
		})
	}
}

var result int

func Benchmark_sumOfLeftLeaves(b *testing.B) {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	var res int
	for n := 0; n < b.N; n++ {
		res = sumOfLeftLeaves(root)
	}
	result = res
}
