package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodNodes(t *testing.T) {
	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	tree2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
			},
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
		{"example1", tree1, 4},
		{"example2", tree2, 3},
		{"example3", tree3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := goodNodes(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := goodNodes2(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}
}
