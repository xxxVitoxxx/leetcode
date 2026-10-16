package main

import "testing"

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    []*ListNode
		expected *ListNode
	}{{
		name: "example1",
		lists: []*ListNode{
			{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
			{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			{
				Val: 2,
				Next: &ListNode{
					Val: 6,
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
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 6,
									},
								},
							},
						},
					},
				},
			},
		},
	}, {
		name:     "example2",
		lists:    nil,
		expected: nil,
	}, {
		name: "example3",
		lists: []*ListNode{
			{
				Val:  6,
				Next: &ListNode{Val: 10},
			},
			{
				Val:  1,
				Next: &ListNode{Val: 9},
			},
			{
				Val: 4,
			},
			{
				Val:  4,
				Next: &ListNode{Val: 6},
			},
			{
				Val: 8,
			},
		},
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 10,
									},
								},
							},
						},
					},
				},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listNodes := make([]*ListNode, len(tt.lists))
			for i := range len(tt.lists) {
				listNodes[i] = cloneHead(tt.lists[i])
			}

			result := mergeKLists(listNodes)

			h1, h2 := result, tt.expected
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
			listNodes := make([]*ListNode, len(tt.lists))
			for i := range len(tt.lists) {
				listNodes[i] = cloneHead(tt.lists[i])
			}

			result := mergeKLists2(listNodes)

			h1, h2 := result, tt.expected
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
