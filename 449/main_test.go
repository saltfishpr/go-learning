package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var c = Constructor()

func TestCodec(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
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
		},
		{
			name: "示例 2",
			args: args{
				root: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := c.serialize(tt.args.root)
			root := c.deserialize(data)
			assert.Equal(t, tt.args.root, root)
		})
	}
}

var (
	result1 string
	result2 *TreeNode

	root = &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
)

func BenchmarkCodec_serialize(b *testing.B) {
	var res string
	for n := 0; n < b.N; n++ {
		res = c.serialize(root)
	}
	result1 = res
}

func BenchmarkCodec_deserialize(b *testing.B) {
	var res *TreeNode
	for n := 0; n < b.N; n++ {
		res = c.deserialize("2,1,3")
	}
	result2 = res
}
