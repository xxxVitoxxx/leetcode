package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumAverageSubtree(t *testing.T) {
	node1 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 6},
		Right: &TreeNode{Val: 1},
	}
	node2 := &TreeNode{
		Val:   0,
		Right: &TreeNode{Val: 1},
	}
	// [7,4,9,1,6,null,11]
	node3 := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 11},
		},
	}

	tests := []struct {
		name   string
		root   *TreeNode
		output float64
	}{
		{"example1", node1, 6},
		{"example2", node2, 1},
		{"example3", node3, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximumAverageSubtree(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximumAverageSubtree2(tt.root)
			assert.Equal(t, tt.output, res)
		})
	}
}
