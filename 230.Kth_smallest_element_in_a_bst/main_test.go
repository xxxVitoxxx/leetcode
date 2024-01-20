package main

import "testing"

func TestKthSmallest(t *testing.T) {
	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	tests := []struct {
		name   string
		root   *TreeNode
		k      int
		output int
	}{
		{"example1", tree1, 1, 1},
		{"example2", tree2, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := kthSmallest(tt.root, tt.k)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := kthSmallest2(tt.root, tt.k)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
