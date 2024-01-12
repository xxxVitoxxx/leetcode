package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestZigZag(t *testing.T) {
	tree1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	tree3 := &TreeNode{
		Val: 1,
	}

	tests := []struct {
		name   string
		root   *TreeNode
		output int
	}{
		{"example1", tree1, 3},
		{"example2", tree2, 4},
		{"example3", tree3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestZigZag(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestZigZag2(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}
}
