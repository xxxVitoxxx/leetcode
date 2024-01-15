package main

import "testing"

func TestIsSameTree(t *testing.T) {
	tree1P := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tree1Q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	tree2P := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	tree2Q := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	tree3P := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	tree3Q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	tests := []struct {
		name   string
		p      *TreeNode
		q      *TreeNode
		output bool
	}{
		{"example1", tree1P, tree1Q, true},
		{"example2", tree2P, tree2Q, false},
		{"example3", tree3P, tree3Q, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSameTree(tt.p, tt.q)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSameTree2(tt.p, tt.q)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
