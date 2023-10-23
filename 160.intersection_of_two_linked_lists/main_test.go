package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	// case1
	intersectNode1 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	headA1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: intersectNode1,
		},
	}
	headB1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: intersectNode1,
			},
		},
	}

	// case2
	intersectNode2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}
	headA2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  1,
				Next: intersectNode2,
			},
		},
	}
	headB2 := &ListNode{
		Val:  3,
		Next: intersectNode2,
	}

	// case3
	headA3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	headB3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
		},
	}

	// case4
	intersectNode4 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 1,
		},
	}
	headA4 := intersectNode4
	headB4 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  10,
						Next: intersectNode4,
					},
				},
			},
		},
	}

	tests := []struct {
		name   string
		headA  *ListNode
		headB  *ListNode
		output *ListNode
	}{
		{"example1", headA1, headB1, intersectNode1},
		{"example2", headA2, headB2, intersectNode2},
		{"example3", headA3, headB3, nil},
		{"exampl4", headA4, headB4, intersectNode4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getIntersectionNode(tt.headA, tt.headB)
			assert.True(t, reflect.DeepEqual(res, tt.output))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getIntersectionNode2(tt.headA, tt.headB)
			assert.True(t, reflect.DeepEqual(res, tt.output))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getIntersectionNode3(tt.headA, tt.headB)
			assert.True(t, reflect.DeepEqual(res, tt.output))
		})
	}
}
