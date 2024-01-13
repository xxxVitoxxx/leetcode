package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	tree2q := &TreeNode{
		Val: 4,
	}
	tree1p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: tree2q,
		},
	}
	tree1q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	tree := &TreeNode{
		Val:   3,
		Left:  tree1p,
		Right: tree1q,
	}

	tree3q := &TreeNode{
		Val: 2,
	}
	tree3 := &TreeNode{
		Val:  1,
		Left: tree3q,
	}

	tests := []struct {
		name   string
		root   *TreeNode
		p      *TreeNode
		q      *TreeNode
		output int
	}{
		{"example1", tree, tree1p, tree1q, 3},
		{"example2", tree, tree1p, tree2q, 5},
		{"example3", tree3, tree3, tree3q, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lowestCommonAncestor(tt.root, tt.p, tt.q)
			assert.Equal(t, tt.output, res.Val)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lowestCommonAncestor2(tt.root, tt.p, tt.q)
			assert.Equal(t, tt.output, res.Val)
		})
	}
}
