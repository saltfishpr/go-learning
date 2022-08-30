package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例 1",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 3,
							Children: []*Node{
								{Val: 5},
								{Val: 6},
							},
						},
						{Val: 2},
						{Val: 4},
					},
				},
			},
			want: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			name: "示例 2",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{Val: 2},
						{
							Val: 3,
							Children: []*Node{
								{Val: 6},
								{
									Val: 7,
									Children: []*Node{
										{
											Val: 11,
											Children: []*Node{
												{Val: 14},
											},
										},
									},
								},
							},
						},
						{
							Val: 4,
							Children: []*Node{
								{
									Val: 8,
									Children: []*Node{
										{Val: 12},
									},
								},
							},
						},
						{
							Val: 5,
							Children: []*Node{
								{
									Val: 9,
									Children: []*Node{
										{Val: 13},
									},
								},
								{Val: 10},
							},
						},
					},
				},
			},
			want: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, levelOrder(tt.args.root))
		})
	}
}

var result [][]int

func Benchmark_levelOrder(b *testing.B) {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{Val: 2},
			{Val: 4},
		},
	}
	var res [][]int
	for n := 0; n < b.N; n++ {
		res = levelOrder(root)
	}
	result = res
}
