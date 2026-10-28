package main

import "testing"

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		subRoot  *TreeNode
		expected bool
	}{{
		name: "example1",
		root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{Val: 5},
		},
		subRoot: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		expected: true,
	}, {
		name: "example2",
		root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{Val: 5},
		},
		subRoot: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
			},
		},
		expected: false,
	}, {
		name: "example3",
		root: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 1},
		},
		subRoot:  &TreeNode{Val: 1},
		expected: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isSubtree(tt.root, tt.subRoot)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isSubtree2(tt.root, tt.subRoot)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
