package main

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		p        *TreeNode
		q        *TreeNode
		expected int
	}{{
		name: "example1",
		root: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
			},
			Right: &TreeNode{
				Val:   9,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 9},
			},
		},
		p:        &TreeNode{Val: 2},
		q:        &TreeNode{Val: 8},
		expected: 6,
	}, {
		name: "example1",
		root: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
			},
			Right: &TreeNode{
				Val:   9,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 9},
			},
		},
		p:        &TreeNode{Val: 2},
		q:        &TreeNode{Val: 4},
		expected: 2,
	}, {
		name: "example1",
		root: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1},
		},
		p:        &TreeNode{Val: 2},
		q:        &TreeNode{Val: 1},
		expected: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lowestCommonAncestor(tt.root, tt.p, tt.q)
			if actual.Val != tt.expected {
				t.Errorf("got: %v, want: %v", actual.Val, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lowestCommonAncestor2(tt.root, tt.p, tt.q)
			if actual.Val != tt.expected {
				t.Errorf("got: %v, want: %v", actual.Val, tt.expected)
			}
		})
	}
}
