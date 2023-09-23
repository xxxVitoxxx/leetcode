package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	// [3,9,20,null,null,15,7]
	input1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	// [1,null,2]
	input2 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}

	tests := []struct {
		name   string
		input  *TreeNode
		output int
	}{
		{"example1", input1, 3},
		{"example2", input2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxDepth(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
