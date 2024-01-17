package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	tree2 := &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	tree3 := &TreeNode{
		Val: 1,
	}

	tests := []struct {
		name   string
		root   *TreeNode
		output [][]int
	}{
		{"example1", tree1, [][]int{{3}, {9, 20}, {15, 7}}},
		{"example2", tree2, [][]int{{9}, {5}, {1, 7}}},
		{"example3", tree3, [][]int{{1}}},
		{"example4", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := levelOrder(tt.root)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := levelOrder2(tt.root)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
