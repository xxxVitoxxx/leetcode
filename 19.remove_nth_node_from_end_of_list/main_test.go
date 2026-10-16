package main

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		n        int
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
		n: 2,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}, {
		name:     "example2",
		head:     &ListNode{Val: 1},
		n:        1,
		expected: nil,
	}, {
		name: "example3",
		head: &ListNode{
			Val:  1,
			Next: &ListNode{Val: 2},
		},
		n:        1,
		expected: &ListNode{Val: 1},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newHead := removeNthFromEnd(cloneHead(tt.head), tt.n)

			h1, h2 := newHead, tt.expected
			for h1 != nil && h2 != nil {
				if h1.Val != h2.Val {
					t.Errorf("value is not equal")
				}

				h1 = h1.Next
				h2 = h2.Next
			}

			if h1 != nil || h2 != nil {
				t.Errorf("some head still have data")
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newHead := removeNthFromEnd2(cloneHead(tt.head), tt.n)

			h1, h2 := newHead, tt.expected
			for h1 != nil && h2 != nil {
				if h1.Val != h2.Val {
					t.Errorf("value is not equal")
				}

				h1 = h1.Next
				h2 = h2.Next
			}

			if h1 != nil || h2 != nil {
				t.Errorf("some head still have data")
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newHead := removeNthFromEnd3(cloneHead(tt.head), tt.n)

			h1, h2 := newHead, tt.expected
			for h1 != nil && h2 != nil {
				if h1.Val != h2.Val {
					t.Errorf("value is not equal")
				}

				h1 = h1.Next
				h2 = h2.Next
			}

			if h1 != nil || h2 != nil {
				t.Errorf("some head still have data")
			}
		})
	}
}

func cloneHead(head *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for node := head; node != nil; node = node.Next {
		curr.Next = &ListNode{Val: node.Val}
		curr = curr.Next
	}

	return dummy.Next
}
