package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 9},
					},
				},
			},
			want: []int{1, 3, 9},
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 3},
		},
		{
			name: "示例 3",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestValues(tt.args.root))
		})
	}
}

var result []int

func Benchmark_largestValues(b *testing.B) {
	var res []int

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 9},
		},
	}

	for n := 0; n < b.N; n++ {
		res = largestValues(root)
	}

	result = res
}
