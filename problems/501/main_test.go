package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMode(t *testing.T) {
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
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 2},
					},
				},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMode(tt.args.root))
		})
	}
}

var result []int

func Benchmark_findMode(b *testing.B) {
	var res []int

	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 2},
		},
	}

	for n := 0; n < b.N; n++ {
		res = findMode(root)
	}

	result = res
}
