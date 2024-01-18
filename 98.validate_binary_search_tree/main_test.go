package main

import (
	"testing"
)

func TestIsValidTree(t *testing.T) {
	tree1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	tree3 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	tree4 := &TreeNode{
		Val: 2147483647,
	}
	tree5 := &TreeNode{
		Val: -2147483648,
	}

	tests := []struct {
		name   string
		root   *TreeNode
		output bool
	}{
		{"example1", tree1, true},
		{"example2", tree2, false},
		{"example3", tree3, false},
		{"example4", tree4, true},
		{"example5", tree5, true},
		{"example6", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isValidBST(tt.root)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isValidBST2(tt.root)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isValidBST3(tt.root)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
