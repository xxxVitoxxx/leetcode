package main

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tree1 := &TreeNode{
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
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	tests := []struct {
		name string
		root *TreeNode
		ans  []int
	}{
		{"example1", tree1, []int{4, 7, 2, 9, 6, 3, 1}},
		{"example2", tree2, []int{2, 3, 1}},
		{"example3", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := invertTree(tt.root)
			ans := getTreeValues(res)
			if !reflect.DeepEqual(ans, tt.ans) {
				t.Errorf("got %v, want %v", ans, tt.ans)
			}
		})
	}
}

func TestInvertTree2(t *testing.T) {
	tree1 := &TreeNode{
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
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	tests := []struct {
		name string
		root *TreeNode
		ans  []int
	}{
		{"example1", tree1, []int{4, 7, 2, 9, 6, 3, 1}},
		{"example2", tree2, []int{2, 3, 1}},
		{"example3", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := invertTree2(tt.root)
			ans := getTreeValues(res)
			if !reflect.DeepEqual(ans, tt.ans) {
				t.Errorf("got %v, want %v", ans, tt.ans)
			}
		})
	}
}

func getTreeValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		ans = append(ans, element.Val)
		if element.Left != nil {
			queue = append(queue, element.Left)
		}
		if element.Right != nil {
			queue = append(queue, element.Right)
		}
	}
	return ans
}
