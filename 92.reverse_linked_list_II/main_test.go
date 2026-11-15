package main

import "testing"

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		left     int
		right    int
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
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		left:  2,
		right: 4,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	}, {
		name: "example2",
		head: &ListNode{
			Val: 5,
		},
		left:  1,
		right: 1,
		expected: &ListNode{
			Val: 5,
		},
	}, {
		name: "example3",
		head: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
		left:  1,
		right: 2,
		expected: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(cloneListNode(tt.head), tt.left, tt.right)
			for want := cloneListNode(tt.expected); want != nil; want, got = want.Next, got.Next {
				if got == nil || want.Val != got.Val {
					t.Errorf("want %d, got %d", want.Val, got.Val)
				}
			}
		})
	}
}

func cloneListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{
		Val:  head.Val,
		Next: cloneListNode(head.Next),
	}
	return newHead
}
