package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView(t *testing.T) {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	tree2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
		},
	}
	tree4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	tests := []struct {
		name   string
		root   *TreeNode
		output []int
	}{
		{"example1", tree1, []int{1, 3, 4}},
		{"example2", tree2, []int{1, 3}},
		{"example3", nil, []int{}},
		{"example4", tree4, []int{1, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := rightSideView(tt.root)
			if res == nil {
				assert.Len(t, tt.output, 0)
			} else {
				assert.True(t, reflect.DeepEqual(tt.output, res))
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := rightSideView2(tt.root)
			if res == nil {
				assert.Len(t, tt.output, 0)
			} else {
				assert.True(t, reflect.DeepEqual(tt.output, res))
			}
		})
	}
}
