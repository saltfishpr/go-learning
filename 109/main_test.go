package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "示例 1",
			args: args{
				head: &ListNode{
					Val: -10,
					Next: &ListNode{
						Val: -3,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 9,
								},
							},
						},
					},
				},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
		{
			name: "示例 2",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sortedListToBST(tt.args.head))
		})
	}
}

var result *TreeNode

func Benchmark_sortedListToBST(b *testing.B) {
	head := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}

	var res *TreeNode
	for n := 0; n < b.N; n++ {
		res = sortedListToBST(head)
	}
	result = res
}
