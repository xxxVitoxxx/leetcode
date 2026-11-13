package main

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{{
		name: "example1",
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		expected: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}, {
		name: "example2",
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		expected: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}, {
		name:     "example3",
		head:     &ListNode{Val: 1},
		expected: &ListNode{Val: 1},
	}, {
		name:     "example4",
		head:     &ListNode{},
		expected: &ListNode{},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(cloneHead(tt.head))
			cloneGot := cloneHead(got)

			for want := cloneHead(tt.expected); want != nil; want, cloneGot = want.Next, cloneGot.Next {
				if cloneGot == nil || cloneGot.Val != want.Val {
					t.Errorf("got %v, want %v", cloneGot, want.Val)
				}
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs2(cloneHead(tt.head))
			cloneGot := cloneHead(got)

			for want := cloneHead(tt.expected); want != nil; want, cloneGot = want.Next, cloneGot.Next {
				if cloneGot == nil || cloneGot.Val != want.Val {
					t.Errorf("got %v, want %v", cloneGot, want.Val)
				}
			}
		})
	}
}

func cloneHead(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{
		Val:  head.Val,
		Next: cloneHead(head.Next),
	}
	return newHead
}
