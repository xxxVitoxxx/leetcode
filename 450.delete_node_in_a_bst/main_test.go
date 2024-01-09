package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	tree1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ans1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	tests := []struct {
		name   string
		root   *TreeNode
		key    int
		output *TreeNode
	}{
		{"example1", tree1, 3, ans1},
		{"exampel2", tree2, 0, tree2},
		{"example3", &TreeNode{}, 1, &TreeNode{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.root
			res := deleteNode(root, tt.key)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.root
			res := deleteNode2(root, tt.key)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
