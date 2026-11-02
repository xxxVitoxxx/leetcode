package main

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{{
		name: "example1",
		root: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		},
		expected: []int{3, 2, 1},
	}, {
		name: "example2",
		root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 9},
				},
			},
		},
		expected: []int{4, 6, 7, 5, 2, 9, 8, 3, 1},
	}, {
		name:     "example3",
		root:     nil,
		expected: nil,
	}, {
		name:     "example4",
		root:     &TreeNode{Val: 1},
		expected: []int{1},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := postorderTraversal(tt.root)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
