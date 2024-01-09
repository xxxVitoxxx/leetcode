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

	tree3 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}
	ans3 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
	}

	tree4 := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 7,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	ans4 := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 9,
		},
	}

	tree5 := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val: 12,
					Left: &TreeNode{
						Val: 8,
						Right: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 13,
					},
				},
			},
		},
	}
	ans5 := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val: 12,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 13,
					},
				},
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
		{"exampel3", tree3, 3, ans3},
		{"exampel4", tree4, 7, ans4},
		{"exampel5", tree5, 7, ans5},
		{"example6", &TreeNode{}, 1, &TreeNode{}},
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
