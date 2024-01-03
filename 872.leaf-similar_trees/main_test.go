package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeafSimilar(t *testing.T) {
	tree1root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	tree1root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	tree2root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tree2root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	tree3root1 := &TreeNode{
		Val: 1,
	}
	tree3root2 := &TreeNode{}

	tests := []struct {
		name   string
		root1  *TreeNode
		root2  *TreeNode
		output bool
	}{
		{"example1", tree1root1, tree1root2, true},
		{"example2", tree2root1, tree2root2, false},
		{"example3", tree3root1, tree3root2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := leafSimilar(tt.root1, tt.root2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := leafSimilar2(tt.root1, tt.root2)
			assert.Equal(t, tt.output, res)
		})
	}
}
