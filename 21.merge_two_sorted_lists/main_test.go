package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{{
		name: "example1",
		list1: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		list2: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
	}, {
		name:     "example2",
		list1:    nil,
		list2:    nil,
		expected: nil,
	}, {
		name:     "example3",
		list1:    nil,
		list2:    &ListNode{Val: 0},
		expected: &ListNode{Val: 0},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := deepCopyListNode(tt.list1)
			list2 := deepCopyListNode(tt.list2)
			actual := mergeTwoLists(list1, list2)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := deepCopyListNode(tt.list1)
			list2 := deepCopyListNode(tt.list2)
			actual := mergeTwoLists2(list1, list2)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}

func deepCopyListNode(node *ListNode) *ListNode {
	if node == nil {
		return node
	}
	newNode := &ListNode{Val: node.Val}
	curr := newNode
	for p := node.Next; p != nil; p = p.Next {
		curr.Next = &ListNode{Val: p.Val}
		curr = curr.Next
	}
	return newNode
}
