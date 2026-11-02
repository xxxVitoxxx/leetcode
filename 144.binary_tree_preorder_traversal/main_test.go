package main

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
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
		expected: []int{1, 2, 3},
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
		expected: []int{1, 2, 4, 5, 6, 7, 3, 8, 9},
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
			actual := preorderTraversal(tt.root)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
