package main

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{{
		name: "example1",
		root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		expected: 3,
	}, {
		name: "example2",
		root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		expected: 1,
	}, {
		name: "example3",
		root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		expected: 4,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diameterOfBinaryTree(tt.root)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
