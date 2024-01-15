package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLevelSum(t *testing.T) {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: -8,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	tree2 := &TreeNode{
		Val: 989,
		Right: &TreeNode{
			Val: 10250,
			Left: &TreeNode{
				Val: 98693,
			},
			Right: &TreeNode{
				Val: -89388,
				Right: &TreeNode{
					Val: -32127,
				},
			},
		},
	}

	tests := []struct {
		name   string
		root   *TreeNode
		output int
	}{
		{"example1", tree1, 2},
		{"example2", tree2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxLevelSum(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxLevelSum2(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}
}
