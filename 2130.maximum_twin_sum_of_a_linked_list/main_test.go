package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSum(t *testing.T) {
	list1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	list2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	list3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 100000,
		},
	}

	tests := []struct {
		name   string
		head   *ListNode
		output int
	}{
		{"exampel1", list1, 6},
		{"exampel2", list2, 7},
		{"exampel3", list3, 100001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pairSum2(tt.head)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pairSum(tt.head)
			assert.Equal(t, tt.output, res)
		})
	}
}
