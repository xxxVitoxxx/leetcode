package main

import "testing"

func TestCountNodes(t *testing.T) {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	tree2 := &TreeNode{
		Val: 1,
	}

	tests := []struct {
		name   string
		root   *TreeNode
		output int
	}{
		{"example1", tree1, 6},
		{"example2", tree2, 1},
		{"example3", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countNodes(tt.root)
			if res != tt.output {
				t.Errorf("got %d, want %d", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countNodes2(tt.root)
			if res != tt.output {
				t.Errorf("got %d, want %d", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countNodes3(tt.root)
			if res != tt.output {
				t.Errorf("got %d, want %d", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countNodes4(tt.root)
			if res != tt.output {
				t.Errorf("got %d, want %d", res, tt.output)
			}
		})
	}
}
