package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFrequentTreeSum(t *testing.T) {
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
					Val:   5,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: -3},
				},
			},
			want: []int{2, -3, 4},
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: -5},
				},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, findFrequentTreeSum(tt.args.root))
		})
	}
}

var result []int

func Benchmark_findFrequentTreeSum(b *testing.B) {
	var res []int

	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: -3},
	}

	for n := 0; n < b.N; n++ {
		res = findFrequentTreeSum(root)
	}

	result = res
}
