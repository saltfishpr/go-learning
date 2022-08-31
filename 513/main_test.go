package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findBottomLeftValue(t *testing.T) {
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
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: 1,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  5,
							Left: &TreeNode{Val: 7},
						},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findBottomLeftValue(tt.args.root))
		})
	}
}

var result int

func Benchmark_findBottomLeftValue(b *testing.B) {
	var res int

	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	for n := 0; n < b.N; n++ {
		res = findBottomLeftValue(root)
	}

	result = res
}
