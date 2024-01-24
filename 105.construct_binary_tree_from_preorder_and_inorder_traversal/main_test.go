package main

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
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
		Val: -1,
	}

	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		output   *TreeNode
	}{
		{"example1", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, tree1},
		{"example2", []int{-1}, []int{-1}, tree2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := buildTree(tt.preorder, tt.inorder)
			if !reflect.DeepEqual(*res, *tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := buildTree2(tt.preorder, tt.inorder)
			if !reflect.DeepEqual(*res, *tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
