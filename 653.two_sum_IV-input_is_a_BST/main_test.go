package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTarget(t *testing.T) {
	treeNode1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}

	treeNode2 := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{Val: 7},
		},
	}

	tests := []struct {
		name   string
		input  *TreeNode
		k      int
		output bool
	}{
		{"example1", treeNode1, 9, true},
		{"example2", treeNode1, 28, false},
		{"example3", treeNode2, 10, true},
		{"example4", treeNode2, 28, false},
		{"example5", &TreeNode{}, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTarget(tt.input, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTarget2(tt.input, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTarget3(tt.input, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTarget4(tt.input, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
