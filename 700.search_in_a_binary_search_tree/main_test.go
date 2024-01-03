package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBST(t *testing.T) {
	output1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tree1 := &TreeNode{
		Val:  4,
		Left: output1,
		Right: &TreeNode{
			Val: 7,
		},
	}

	output2 := &TreeNode{}
	tree2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	tests := []struct {
		name   string
		root   *TreeNode
		val    int
		output *TreeNode
	}{
		{"example1", tree1, 2, output1},
		{"example2", tree2, 5, output2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := searchBST(tt.root, tt.val)
			if res != nil {
				assert.True(t, reflect.DeepEqual(tt.output, res))
			} else {
				assert.Nil(t, res)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := searchBST2(tt.root, tt.val)
			if res != nil {
				assert.True(t, reflect.DeepEqual(tt.output, res))
			} else {
				assert.Nil(t, res)
			}
		})
	}
}
